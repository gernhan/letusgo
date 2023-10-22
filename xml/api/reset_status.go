package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gernhan/xml/db"
	"github.com/gernhan/xml/repositories"
	"github.com/gernhan/xml/utils"
)

type ResetStatusRequest struct {
	Status  int   `json:"status"`
	BillRunId int64 `json:"billRun"`
	BatchSize int `json:"batchSize"`
}

func ResetStatusesHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody ResetStatusRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Get the batchSize from the request body or use a default value (e.g., 1000).
	batchSize := requestBody.BatchSize
	if batchSize <= 0 {
		batchSize = 500
	}

	// Call the function to perform the batch update.
	err := batchUpdateStatic(requestBody, batchSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message or HTTP status code.
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Updated statuses in the database")
}

// batchUpdateStatic performs batch update on the list of users.
func batchUpdateStatic(requestBody ResetStatusRequest, batchUpdateSize int) error {
	startTime := time.Now()
	var wg sync.WaitGroup

	minMax, err := repositories.FindMinMaxBillId(requestBody.BillRunId, 1)
	if err != nil {
		return fmt.Errorf("cannot query range of bill run id, body %v", requestBody)
	}

	log.Printf("Range of id: %v", minMax)

	partitions, _ := utils.DoPartition(minMax.Min, minMax.Max, 10)
	log.Printf("Partitions %v", partitions)

	var e error
	for i := 0; i < len(partitions); i += batchUpdateSize {
		wg.Add(1)
		finalI := i
		go func() {
			err := batchUpdate(partitions[finalI], &wg, requestBody.BillRunId)
			if err != nil {
				log.Printf("Caught error: %v", err)
				e = err
			}
		}()
	}

	wg.Wait()

	if e != nil {
		return e
	}

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	log.Printf("batchUpdateStatic -> in %v", executionTime)
	return nil
}

func batchUpdate(partition utils.Partition, wg *sync.WaitGroup, billRunId int64) error {
	defer wg.Done()
  db := db.GetPool()

	query := buildUpdateQuery(partition, billRunId)

	// Execute the batch insert query using the connection pool.
	conn, err := db.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}

func buildUpdateQuery(partition utils.Partition, billRunId int64) string {
	var query strings.Builder

	query.WriteString(fmt.Sprintf("UPDATE d_bills db SET pdf_status = %d FROM d_billruns dbr WHERE db.id < %d and db.id >= %d and dbr.id = %d", 0, partition.Max, partition.Min, billRunId))

	return query.String()
}
