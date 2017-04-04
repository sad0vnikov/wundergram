package wservice

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/sad0vnikov/wundergram/wunderlist/wobjects"
)

func TestFilterTasksLessThanDate(t *testing.T) {

	testDate := time.Date(2017, time.January, 2, 0, 0, 0, 0, time.UTC)

	expectedTask := makeTestTask("test01", "2016-01-03")
	log.Print(expectedTask)

	testTasks := []wobjects.Task{
		expectedTask,
		makeTestTask("test02", "2017-03-01"),
		makeTestTask("test03", "2017-03-05"),
	}

	filteredTasks := filterTasksLessThanDate(testDate, testTasks)
	if len(filteredTasks) < 1 {
		t.Error("Got an empty filtered tasks list")
	}

	if len(filteredTasks) > 1 {
		t.Errorf("Too many tasks filtered: %#v", filteredTasks)
	}

	if filteredTasks[0].Title != expectedTask.Title {
		t.Errorf("Got wrong task in filter: extected task %v, got %v", expectedTask, filteredTasks[0])
	}

}

func makeTestTask(title string, dueDate string) wobjects.Task {
	return wobjects.Task{
		ID:          rand.Int(),
		AssigneeID:  rand.Int(),
		CreatedAt:   "2013-08-30T08:36:13.273Z",
		CreatedByID: rand.Int(),
		DueDate:     dueDate,
		ListID:      1,
		Revision:    1,
		Starred:     false,
		Title:       title,
	}
}
