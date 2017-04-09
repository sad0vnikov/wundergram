package util

import (
	"time"

	"github.com/bradfitz/latlong"
	"github.com/sad0vnikov/wundergram/logger"
)

//GetDayStart gets a date and returns a time for day start
func GetDayStart(currentTime time.Time, tz *time.Location) time.Time {
	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, tz)
}

//GetTimezoneByCoord returns user's time.Location by latitude and longtitude
func GetTimezoneByCoord(lat, long float64) *time.Location {
	zoneName := latlong.LookupZoneName(lat, long)
	loc, err := time.LoadLocation(zoneName)
	if err != nil {
		logger.Get("main").Errorf("error detecting user timezone: lat = %v, long = %v, got zone name %v, could't parse. error: %v", lat, long, zoneName, err)
		loc = time.UTC
	}
	return loc
}
