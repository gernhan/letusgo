package processor

import (
	"github.com/gernhan/xml/concurrent"
	"sync"
)

type DataBatchProcessor struct {
	buffer          []interface{}
	bufferThreshold int
	handler         func([]interface{}) (interface{}, error)
	lock            *sync.Mutex
	nextProcessor   *DataBatchProcessor
	results         *concurrent.List
}

func NewDataBatchProcessor(
	bufferThreshold int,
	handler func([]interface{}) (interface{}, error),
	storingResult bool,
	nextProcessor *DataBatchProcessor,
) *DataBatchProcessor {
	var results *concurrent.List
	if storingResult {
		results = concurrent.NewConcurrentList()
	}
	return &DataBatchProcessor{
		buffer:          make([]interface{}, 0),
		bufferThreshold: bufferThreshold,
		handler:         handler,
		lock:            &sync.Mutex{},
		results:         results,
		nextProcessor:   nextProcessor,
	}
}

func (d *DataBatchProcessor) Handle(data interface{}) (interface{}, error) {
	d.lock.Lock()
	d.buffer = append(d.buffer, data)
	if len(d.buffer) >= d.bufferThreshold {
		dataInput := make([]interface{}, len(d.buffer))
		copy(dataInput, d.buffer)
		d.buffer = d.buffer[:0]
		d.lock.Unlock()
		result, err := d.handler(dataInput)
		if err == nil {
			if d.results != nil {
				d.results.Add(result)
			}
			if d.nextProcessor != nil {
				_, err = d.nextProcessor.Handle(result)
			}
		}
		return result, err
	}
	d.lock.Unlock()
	return nil, nil
}

func (d *DataBatchProcessor) HandleLastBatch() (interface{}, error) {
	d.lock.Lock()
	if len(d.buffer) > 0 {
		dataInput := make([]interface{}, len(d.buffer))
		copy(dataInput, d.buffer)
		d.buffer = d.buffer[:0]
		d.lock.Unlock()
		result, err := d.handler(dataInput)
		if d.results != nil {
			d.results.Add(result)
		}
		if d.nextProcessor != nil {
			_, err = d.nextProcessor.Handle(result)
		}
		return result, err
	}
	d.lock.Unlock()
	return nil, nil
}

func (d *DataBatchProcessor) Results() *concurrent.List {
	return d.results
}
