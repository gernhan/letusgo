package utils

import (
	"time"
)

var dateFormatter = "02.01.2006"
var timeFormatter = "15:04:05"

func formatDate(date time.Time) string {
	if date.IsZero() {
		return ""
	}
	return date.Format(dateFormatter)
}

func formatTime(dateTime time.Time) string {
	return dateTime.Format(timeFormatter)
}
