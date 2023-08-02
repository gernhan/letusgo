package processor

import (
	"log"
	"sync/atomic"
	"testing"

	cc "github.com/gernhan/template/concurrent"
	"github.com/gernhan/template/tools"
)

func TestDataBatchProcessor(t *testing.T) {
	bufferThreshold := 10
	collector := *cc.NewConcurrentList()
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
	futures := cc.EmptyMultiFutures()
	for i := 0; i < nThreads; i++ {
		futures.AddFuture(cc.RunAsync(func() error {
			_, err := dataBatchProcessor.handle(ig.GenerateString(4))
			return err
		}))
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
	collector := *cc.NewConcurrentList()
	ig := tools.NewInputGenerator()

	dataBatchProcessor := NewDataBatchProcessor(
		bufferThreshold,
		func(dataCollection []interface{}) (interface{}, error) {
			return cc.SupplyAsync(func() (interface{}, error) {
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

	multiFutures := cc.EmptyMultiFutures()
	for i := 0; i < items; i++ {
		f := cc.RunAsync(func() error {
			_, err := dataBatchProcessor.handle(ig.GenerateString(3))
			return err
		})
		multiFutures.AddFuture(f)
	}

	result := multiFutures.Result()
	if !result.IsSucceed {
		err := result.Failures[0].Err
		t.Errorf("Caught error %v", err)
	}
	_, err := dataBatchProcessor.handleLastBatch()
	if err != nil {
		t.Errorf("Caught error %v", err)
	}

	for value := range dataBatchProcessor.Results().Values() {
		future := value.(*cc.Future)
		future.Wait()
		log.Printf("future %v, data %v", future, future.Data())
	}

	if collector.Size() != items {
		t.Errorf("Expected list size %d, actual: %v", items, collector.Size())
	}
}
