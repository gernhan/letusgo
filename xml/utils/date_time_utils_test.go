package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestDateTime(t *testing.T) {
	// Example usage of DateFormatter functions
	calendar := time.Date(2023, time.August, 2, 12, 30, 0, 0, time.UTC)

	fmt.Println(formatDate(calendar)) // Output: 02.08.2023
	fmt.Println(formatTime(calendar)) // Output: 12:30:00
}
