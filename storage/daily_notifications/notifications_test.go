package daily_notifications

import "testing"

func TestParsingStringTime(t *testing.T) {
	assertParsedTimeError(t, "00")
	assertParsedTimeError(t, "")
	assertParsedTime(t, "00:00", 0)
	assertParsedTime(t, "00:01", 60)
	assertParsedTime(t, "01:00", 3600)
	assertParsedTime(t, "23:59", 23*60*60+59*60)
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
