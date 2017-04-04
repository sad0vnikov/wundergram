package wunderlist

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/sad0vnikov/wundergram/wunderlist/wobjects"
)

const getAccessTokenLink = "https://www.wunderlist.com/oauth/access_token"

const apiUrl = "https://a.wunderlist.com/api/v1/"

func makeJSONRequest(url string, params map[string]string) (map[string]string, error) {

	jsonToSend, err := json.Marshal(params)

	log.Printf("making json request to %v: %v", url, string(jsonToSend))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonToSend))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := map[string]string{}

	json.NewDecoder(resp.Body).Decode(&response)
	log.Printf("got response: %v", response)

	return response, nil
}

//makeAPIRequest makes a GET request to the given REST resource
func makeAPIRequest(resource, userToken string) ([]byte, error) {
	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodGet, apiUrl+resource, nil)
	req.Header.Set("X-Access-Token", userToken)
	req.Header.Set("X-Client-Id", clientID)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	jsonBytes := []byte{}

	req.Body.Read(jsonBytes)

	return jsonBytes, nil
}

//GetUserAccessToken returns OAuth access token
func GetUserAccessToken(authCode string) (string, error) {
	params := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          authCode,
	}

	response, err := makeJSONRequest(getAccessTokenLink, params)
	if err != nil {
		return "", err
	}

	return response["access_token"], nil
}

//GetLists returns an array of user's tasks lists
func GetLists(userToken string) ([]wobjects.List, error) {

	jsonResult, err := makeAPIRequest("/lists", userToken)
	if err != nil {
		return nil, err
	}

	lists, err := wobjects.ListArrayFromJSON(jsonResult)
	if err != nil {
		return lists, err
	}

	return lists, nil
}

//GetTasks returns all the tasks dedicated to user
func GetTasks(listID int, userToken string) ([]wobjects.Task, error) {

	jsonResult, err := makeAPIRequest("/tasks?list_id="+strconv.Itoa(listID), userToken)
	if err != nil {
		return nil, err
	}

	userTasks, err := wobjects.TaskArrayFromJSON(jsonResult)
	return userTasks, nil
}

//GetAllTasks returns all tasks list for user
func GetAllTasks(userToken string) ([]wobjects.Task, error) {
	lists, err := GetLists(userToken)
	if err != nil {
		return nil, err
	}

	tasks := []wobjects.Task{}
	for _, l := range lists {
		listTasks, err := GetTasks(l.ID, userToken)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, listTasks...)
	}

	return tasks, nil
}
