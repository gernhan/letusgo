package utils

import (
	"fmt"
	"testing"
)

func TestPartitions(t *testing.T) {
	min := int64(1)
	max := int64(100)
	partitions := int64(5)

	result, err := DoPartition(min, max, partitions)
	if err != nil {
		// Handle the error
	}

	// Print the result
	for _, part := range result {
		fmt.Printf("Partition %d: [%d, %d]\n", part.PartitionNumber, part.Min, part.Max)
	}
}
