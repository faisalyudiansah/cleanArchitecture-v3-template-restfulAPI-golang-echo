package util

import "time"

// YYYY-MM-DD
func DateFormatter(date string) string {
	newFormatDate, _ := time.Parse("2006-01-02T15:04:05-07:00", date)
	return newFormatDate.Format("2006-01-02")
}
