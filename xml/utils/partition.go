package utils

import (
	"errors"
)

// Partition represents a range partition.
type Partition struct {
	Min             int64
	Max             int64
	PartitionNumber int
}

func DoPartition(min, max, partitions int64) ([]Partition, error) {
	if max < min || partitions <= 0 {
		return nil, errors.New("invalid input")
	}

	gap := (max - min + 1) / partitions
	if gap == 0 {
		return []Partition{{Min: min, Max: max, PartitionNumber: 1}}, nil
	}

	var result []Partition
	separator := min
	partitionNumber := 1

	for separator+gap < max {
		result = append(result, Partition{Min: separator, Max: separator + gap - 1, PartitionNumber: partitionNumber})
		partitionNumber++
		separator += gap
	}

	result = append(result, Partition{Min: separator, Max: max, PartitionNumber: partitionNumber})

	return result, nil
}
