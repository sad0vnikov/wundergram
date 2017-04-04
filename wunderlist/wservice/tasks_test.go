package wservice

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/sad0vnikov/wundergram/wunderlist/wobjects"
)

func TestFilteringTasksForToday(t *testing.T) {
	nowtime := time.Now()
	today := fmt.Sprintf("%d-%d-%d", nowtime.Year(), nowtime.Month(), nowtime.Day())

	todayTask := makeTestTask("test03", today)
	testTasks := []wobjects.Task{
		makeTestTask("test01", "2017-01-02"),
		makeTestTask("test02", "2017-03-01"),
		todayTask,
	}

	filteredTasks := filterTasksForToday(testTasks)
	if len(filteredTasks) != 1 {
		t.Error("Too many tasks filtered")
	}

	if filteredTasks[0].Title != todayTask.Title {
		t.Errorf("Got wrong task in filter: extected task %v, got %v", todayTask, filteredTasks[0])
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
