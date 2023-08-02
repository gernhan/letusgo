package processor

import (
	"log"
	"sync/atomic"
	"testing"

	"github.com/gernhan/template/concurrent"
	"github.com/gernhan/template/tools"
)

func TestDataBatchProcessor(t *testing.T) {
	bufferThreshold := 10
	collector := *concurrent.NewConcurrentList()
	ig := tools.NewInputGenerator()
	counter := int64(0)

	dataBatchProcessor := NewDataBatchProcessor(
		bufferThreshold,
		func(dataCollection []interface{}) (interface{}, error) {
			log.Printf("Handling Batch ... %v\n", dataCollection)
			for _, data := range dataCollection {
				collector.Add(data)
				atomic.AddInt64(&counter, 1)
			}
			v := []string{"Handled Batch"}
			log.Println(v)
			return v, nil
		},
		false,
		nil,
	)

	nThreads := 103
	tp := concurrent.NewThreadPool(3)
	futures := concurrent.EmptyMultiFutures()
	for i := 0; i < nThreads; i++ {
		futures.AddFuture(concurrent.RunAsyncWithPool(func() error {
			_, err := dataBatchProcessor.handle(ig.GenerateString(4))
			return err
		}, tp))
	}

	result := futures.Result()
	if !result.IsSucceed {
		err := result.Failures[0].Err
		t.Errorf("Caught error %v", err)
	}
	_, err := dataBatchProcessor.handleLastBatch()
	if err != nil {
		t.Errorf("Caught error %v", err)
	}

	log.Printf("Actions: %v", counter)
	if collector.Size() != nThreads {
		t.Errorf("Expected list size %d, actual: %v", nThreads, collector.Size())
	}
}

func TestDataBatchProcessorIfHandlerIsConcurrent(t *testing.T) {
	bufferThreshold := 10
	collector := *concurrent.NewConcurrentList()
	ig := tools.NewInputGenerator()

	tp := concurrent.NewThreadPool(3)
	dataBatchProcessor := NewDataBatchProcessor(
		bufferThreshold,
		func(dataCollection []interface{}) (interface{}, error) {
			return tp.Submit(func() (interface{}, error) {
				for _, data := range dataCollection {
					collector.Add(data)
				}
				return "Done", nil
			}), nil
		},
		true,
		nil,
	)

	items := 102

	handleFlows := concurrent.EmptyMultiFutures()
	for i := 0; i < items; i++ {
		f := concurrent.RunAsync(func() error {
			_, err := dataBatchProcessor.handle(ig.GenerateString(3))
			return err
		})
		handleFlows.AddFuture(f)
	}

	f := handleFlows.ToFuture()
	f.Wait()
	result := f.Data().(*concurrent.AllResults)
	if !result.IsSucceed {
		err := result.Failures[0].Err
		t.Errorf("Caught error %v", err)
	}

	_, err := dataBatchProcessor.handleLastBatch()
	if err != nil {
		t.Errorf("Caught error %v", err)
	}

	for value := range dataBatchProcessor.Results().Values() {
		future := value.(*concurrent.Future)
		future.Wait()
		log.Printf("future %v, data %v", future, future.Data())
	}

	if collector.Size() != items {
		t.Errorf("Expected list size %d, actual: %v", items, collector.Size())
	}
}
