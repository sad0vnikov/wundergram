package daily_notifications

import (
	"testing"
	"time"
)

func TestParsingStringTime(t *testing.T) {
	assertParsedTimeError(t, "00")
	assertParsedTimeError(t, "")
	assertParsedTime(t, "00:00", 0)
	assertParsedTime(t, "00:01", 60)
	assertParsedTime(t, "01:00", 3600)
	assertParsedTime(t, "23:59", 23*60*60+59*60)
}

func TestCheckingNotificationWasSentToday(t *testing.T) {
	notification := DailyNotificationConfig{
		UserID:                0,
		NotificationTimestamp: 3600 * 12,
		LastTimeActivated:     0,
	}

	curTime := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	wasSent := notification.CheckWasSentToday(curTime)
	if wasSent != false {
		t.Error("wrong CheckWasSentToday() value, expected false")
	}

	notification.LastTimeActivated = curTime.Unix()
	wasSent = notification.CheckWasSentToday(curTime)
	if wasSent != true {
		t.Error("wrong CheckWasSentToday() value, expected true")
	}

	notification.LastTimeActivated = curTime.Add(time.Second).Unix()
	wasSent = notification.CheckWasSentToday(curTime)
	if wasSent != true {
		t.Error("wrong CheckWasSentToday() value, expected true")
	}

}

func TestCheckingNotificationTimeToSend(t *testing.T) {
	notification := DailyNotificationConfig{
		UserID:                0,
		NotificationTimestamp: 3600 * 12,
		LastTimeActivated:     0,
	}
	curTime := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	timeToSend := notification.CheckIsTimeToSend(curTime)
	if timeToSend != false {
		t.Error("wrong CheckIsTimeToSend() value, expected false")
	}

	curTime = time.Date(2015, 1, 1, 12, 0, 0, 0, time.UTC)
	timeToSend = notification.CheckIsTimeToSend(curTime)
	if timeToSend != true {
		t.Error("wrong CheckIsTimeToSend() value, expected true")
	}

	curTime = time.Date(2015, 1, 1, 23, 59, 59, 59, time.UTC)
	timeToSend = notification.CheckIsTimeToSend(curTime)
	if timeToSend != true {
		t.Error("wrong CheckIsTimeToSend() value, expected true")
	}
}

func assertParsedTime(t *testing.T, testTime string, expectedResult int) {
	parseResult, err := stringTimeToDayOffset(testTime)

	if err != nil {
		t.Errorf("error parsing string time '%v': error %v", testTime, err)
	}

	if parseResult != expectedResult {
		t.Errorf("error parsing string time '%v': got result %v, expected %v", testTime, parseResult, expectedResult)
	}
}

func assertParsedTimeError(t *testing.T, testTime string) {
	parseResult, err := stringTimeToDayOffset(testTime)
	if err == nil {
		t.Errorf("error testing parsing time '%v': expected error, got result %v", testTime, parseResult)
	}
}
