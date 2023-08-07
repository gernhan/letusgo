package schedule_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/gernhan/xml/concurrent"
	"github.com/gernhan/xml/concurrent/schedule"
)

func TestSchedule(t *testing.T) {
	schedule.Init()
	mtp := concurrent.EmptyMultiFutures()

	mtp.AddFuture(schedule.Schedule(func() (interface{}, error) {
		log.Printf("Scheduled %v", 1)
		return fmt.Sprintf("Item %v", 1), nil
	}, 2*time.Second))

	mtp.AddFuture(schedule.Schedule(func() (interface{}, error) {
		log.Printf("Scheduled %v", 2)
		return fmt.Sprintf("Item %v", 2), nil
	}, 1*time.Second))

	result := mtp.Result()
	if !result.IsSucceed {
		err := result.Failures[0].Err
		t.Errorf("Caught error %v", err)
	}

	for _, item := range mtp.Result().Successes {
		value := item.Data.(string)
		log.Printf("Data %v", value)
	}
}
