package wobjects

import (
	"encoding/json"
)

//Task is a Wunderlist Task
type Task struct {
	ID          int    `json:"id"`
	AssigneeID  int    `json:"assignee_id"`
	CreatedAt   string `json:"created_at"`
	CreatedByID int    `json:"created_by_id"`
	DueDate     string `json:"due_date"`
	ListID      int    `json:"list_id"`
	Revision    int    `json:"revision"`
	Starred     bool   `json:"starred"`
	Title       string `json:"title"`
}

//TaskFromJSON parses a JSON and returns Task object
func TaskFromJSON(j []byte) (Task, error) {
	var task Task
	err := json.Unmarshal(j, task)
	if err != nil {
		return task, err
	}

	return task, nil
}

//TaskArrayFromJSON parses a JSON and returns an array of Tasks
func TaskArrayFromJSON(j []byte) ([]Task, error) {
	var task []Task
	err := json.Unmarshal(j, task)
	if err != nil {
		return task, err
	}

	return task, nil
}
