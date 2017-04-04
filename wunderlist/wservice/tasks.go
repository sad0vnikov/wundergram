package wservice

import (
	"fmt"
	"time"

	"github.com/sad0vnikov/wundergram/storage/tokens"
	"github.com/sad0vnikov/wundergram/wunderlist"
	"github.com/sad0vnikov/wundergram/wunderlist/wobjects"
)

//GetUserTodayTasks returns users's today tasks
func GetUserTodayTasks(userID int) ([]wobjects.Task, error) {
	userToken, err := tokens.Get(userID)
	if err != nil {
		return nil, err
	}

	userTasks, err := wunderlist.GetAllTasks(userToken)
	if err != nil {
		return nil, err
	}
	todayTasks := filterTasksForToday(userTasks)

	return todayTasks, nil
}

func filterTasksForToday(tasks []wobjects.Task) []wobjects.Task {
	todayTasks := make([]wobjects.Task, 0)
	t := time.Now()
	todayDate := fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())

	for _, task := range tasks {
		if task.DueDate == todayDate {
			todayTasks = append(todayTasks, task)
		}
	}

	return todayTasks
}
