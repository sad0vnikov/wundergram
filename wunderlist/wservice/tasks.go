package wservice

import (
	"time"

	"github.com/sad0vnikov/wundergram/logger"
	"github.com/sad0vnikov/wundergram/storage/tokens"
	"github.com/sad0vnikov/wundergram/wunderlist"
	"github.com/sad0vnikov/wundergram/wunderlist/wobjects"
)

//GetUserTodayTasks returns users's today tasks
func GetUserTodayTasks(userID int) ([]wobjects.Task, error) {
	userToken, err := tokens.Get(userID)
	if err != nil {
		logger.Get("main").Error(err)
		return nil, err
	}

	userTasks, err := wunderlist.GetAllTasks(userToken)
	if err != nil {
		logger.Get("main").Error(err)
		return nil, err
	}
	todayTasks := filterTasksForToday(userTasks)

	return todayTasks, nil
}

func filterTasksLessThanDate(date time.Time, tasks []wobjects.Task) []wobjects.Task {
	filteredTasks := make([]wobjects.Task, 0)

	for _, task := range tasks {
		parsedTime, err := time.Parse("2006-01-02", task.DueDate)
		if err == nil && parsedTime.Unix() < date.Unix() {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks
}

func filterTasksForToday(tasks []wobjects.Task) []wobjects.Task {
	tomorrow := time.Now().AddDate(0, 0, 1)
	return filterTasksLessThanDate(tomorrow, tasks)
}
