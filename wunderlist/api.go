package wunderlist

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const getAccessTokenLink = "https://www.wunderlist.com/oauth/access_token"

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
