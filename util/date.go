package util

import "time"

//GetDayStart gets a date and returns a time for day start
func GetDayStart(currentTime time.Time) time.Time {
	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.Local)
}
