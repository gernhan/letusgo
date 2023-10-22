package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gernhan/xml/concurrent/atomics"
	"github.com/gernhan/xml/models"
	"github.com/gorilla/mux"

	"github.com/gernhan/xml/entities/views"
	processor "github.com/gernhan/xml/processors"
	sftputils "github.com/gernhan/xml/sftp"
	"github.com/gernhan/xml/utils"
	xmlhandlers "github.com/gernhan/xml/xml/handlers"

	"github.com/gernhan/xml/concurrent"
	"github.com/gernhan/xml/db"
	"github.com/gernhan/xml/pool"
	"github.com/gernhan/xml/repositories"
)

type XmlGenerationRequest struct {
	BillRunId     int64 `json:"billRunId"`
	DbConnections int64 `json:"dbConnections"`
	BatchInvoices int   `json:"batchInvoices"`
	BatchFiles    int   `json:"batchFiles"`
}

func XmlHandler(w http.ResponseWriter, r *http.Request) {
	requestBody, hasError := createXmlGenerationRequest(r, w)
	if hasError {
		return
	}

	concurrent.RunAsyncWithPool(func() error {
		err := concurrentProcessing(requestBody)
		if err != nil {
			return err
		}
		return nil
	}, pool.GetPools().BillRunPool)

	// Respond with a success message or HTTP status code.
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Finished")
}

func createXmlGenerationRequest(r *http.Request, w http.ResponseWriter) (XmlGenerationRequest, bool) {
	var requestBody XmlGenerationRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return XmlGenerationRequest{}, false
	}

	vars := mux.Vars(r)
	num, err := strconv.ParseInt(vars["billRunId"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid bill run id format", http.StatusBadRequest)
		return XmlGenerationRequest{}, false
	}
	requestBody.BillRunId = num

	if requestBody.DbConnections <= 0 {
		requestBody.DbConnections = 10
	}

	if requestBody.BatchInvoices <= 0 {
		requestBody.BatchInvoices = 100
	}

	if requestBody.BatchFiles <= 0 {
		requestBody.BatchFiles = 100
	}
	return requestBody, true
}

func concurrentProcessing(requestBody XmlGenerationRequest) error {
	minMax, err := repositories.FindMinMaxBillId(requestBody.BillRunId, 0)
	if err != nil {
		return fmt.Errorf("cannot query range of bill run id, body %v", requestBody)
	}
	globalConfig := &models.XmlGenerationGlobalConfiguration{
		BillRunId:     requestBody.BillRunId,
		BatchInvoices: requestBody.BatchInvoices,
		BatchFiles:    requestBody.BatchFiles,

		FileCount:         atomics.NewAtomicInteger(),
		TarCount:          atomics.NewAtomicInteger(),
		UploadedFileCount: concurrent.NewConcurrentMap(),

		TimePrefix: time.Now().Format("2006_01_02_15_04_05"),
	}

	log.Printf("Range of id: %v", minMax)

	partitions, _ := utils.DoPartition(minMax.Min, minMax.Max, requestBody.DbConnections)
	log.Printf("Partitions %v", partitions)

	startTime := time.Now()
	mtp := concurrent.EmptyMultiFutures()
	for i := 0; i < len(partitions); i++ {
		partition := partitions[i]
		config := &models.XmlGenerationConfiguration{
			MaxID: partition.Max,
			MinID: partition.Min,

			GlobalConfig: globalConfig,
		}
		log.Printf("Run with partition %v", partition)

		mtp.AddFuture(concurrent.RunAsyncWithPool(func() error {
			return process(config)
		}, pool.GetPools().DbPool))
	}
	mtp.Wait()
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Time consumed: %v\n", executionTime)
	speed := minMax.Total * 1000 * 3600 / executionTime.Milliseconds()
	fmt.Printf("Speed: %v\n", speed)
	fmt.Printf("Uploaded (%v Files and %v Invoices)\n", globalConfig.UploadedFileCount.Size(), minMax.Total)
	return nil
}

func processDBOnly(config *models.XmlGenerationConfiguration) error {
	billsToXml := processor.NewDataBatchProcessor(
		100,
		func(obj []interface{}) (interface{}, error) {
			bills := make([]views.VExportingBillsV3, 0)
			for _, item := range obj {
				bills = append(bills, item.(views.VExportingBillsV3))
			}

			return concurrent.SupplyAsyncWithPool(func() (interface{}, error) {
				data, err := xmlhandlers.PrepaidHandler.PrepareXmlData(bills)
				if err != nil {
					return nil, err
				}

				fileCount := config.GlobalConfig.FileCount.IncrementAndGet()
				log.Printf(strconv.FormatInt(fileCount, 10))
				fileName := "file_" + strconv.FormatInt(fileCount, 10)
				return &sftputils.FileModel{
					FileName:    fileName,
					FileContent: data,
				}, nil
			}, pool.GetPools().ProcessingPool), nil

		},
		false,
		nil,
	)
	p := db.GetPool()

	billCh, errCh := repositories.FindByBillRunV3(p, config.GlobalConfig.BillRunId, config.MinID, config.MaxID, 0)
	go func() {
		if err := <-errCh; err != nil {
			log.Printf("Caught error %v", err)
		}
	}()

	for bill := range billCh {
		_, err := billsToXml.Handle(bill)
		if err != nil {
			return nil
		}
	}

	_, _ = billsToXml.HandleLastBatch()
	return nil
}

func process(config *models.XmlGenerationConfiguration) error {
	uploadBatchProcessor, billsToXml := createProcessors(config)
	p := db.GetPool()

	billCh, errCh := repositories.FindByBillRunV3(p, config.GlobalConfig.BillRunId, config.MinID, config.MaxID, 0)
	go func() {
		if err := <-errCh; err != nil {
			log.Printf("Caught error %v", err)
		}
	}()

	for bill := range billCh {
		_, err := billsToXml.Handle(bill)
		if err != nil {
			return nil
		}
	}

	_, _ = billsToXml.HandleLastBatch()
	_, _ = uploadBatchProcessor.HandleLastBatch()

	mtp := concurrent.EmptyMultiFutures()
	for value := range uploadBatchProcessor.Results().Values() {
		mtp.AddFuture(value.(*concurrent.Future))
	}

	result := mtp.Result()
	if !result.IsSucceed {
		err := result.Failures[0].Err
		return err
	}
	return nil
}

func createProcessors(config *models.XmlGenerationConfiguration) (*processor.DataBatchProcessor, *processor.DataBatchProcessor) {
	uploadBatchProcessor := processor.NewDataBatchProcessor(
		100,
		func(obj []interface{}) (interface{}, error) {
			mtp := concurrent.EmptyMultiFutures()
			for _, item := range obj {
				mtp.AddFuture(item.(*concurrent.Future))
			}

			result := mtp.Result()
			if !result.IsSucceed {
				err := result.Failures[0].Err
				return nil, err
			}

			fileModels := make([]*sftputils.FileModel, 0)
			for _, item := range mtp.Result().Successes {
				model := item.Data.(*sftputils.FileModel)
				fileModels = append(fileModels, model)
			}

			tarCount := config.GlobalConfig.TarCount.IncrementAndGet()
			fileName := config.GlobalConfig.TimePrefix + "_" +
				strconv.FormatInt(config.GlobalConfig.BillRunId, 10) + "_" +
				strconv.FormatInt(tarCount, 10)
			return concurrent.RunAsyncWithPool(func() error {
				err := sftputils.UploadTarGzV3(
					sftputils.ConnectionConfig{
						Host:        "34.89.153.28",
						Port:        22,
						User:        "compax",
						Password:    "compax",
						PrivateKey:  nil,
						KeyPassword: nil,
					},
					"/upload/test7",
					fileName,
					fileModels,
					config,
				)

				if err != nil {
					return err
				}

				return nil
			}, pool.GetPools().UploadingPool), nil
		},
		true,
		nil,
	)

	billsToXml := processor.NewDataBatchProcessor(
		100,
		func(obj []interface{}) (interface{}, error) {
			bills := make([]views.VExportingBillsV3, 0)
			for _, item := range obj {
				bills = append(bills, item.(views.VExportingBillsV3))
			}

			return concurrent.SupplyAsyncWithPool(func() (interface{}, error) {
				data, err := xmlhandlers.PrepaidHandler.PrepareXmlData(bills)
				if err != nil {
					return nil, err
				}

				fileCount := strconv.FormatInt(config.GlobalConfig.FileCount.IncrementAndGet(), 10)

				builder := strings.Builder{}
				builder.WriteString("file_")
				builder.WriteString(fileCount)
				builder.WriteString(".xml")
				fileName := builder.String()

				return &sftputils.FileModel{
					FileName:    fileName,
					FileContent: data,
				}, nil
			}, pool.GetPools().ProcessingPool), nil

		},
		false,
		uploadBatchProcessor,
	)
	return uploadBatchProcessor, billsToXml
}
