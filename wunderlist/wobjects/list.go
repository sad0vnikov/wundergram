package wobjects

import "encoding/json"

//List is a list of tasks
type List struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Revision  int    `json:"revision"`
}

//ListFromJSON parses given JSON and returns a List
func ListFromJSON(j []byte) (List, error) {
	var list List
	err := json.Unmarshal(j, list)
	if err != nil {
		return list, err
	}

	return list, nil
}

//ListArrayFromJSON parses given JSON and returns an array of List
func ListArrayFromJSON(j []byte) ([]List, error) {
	var list []List
	err := json.Unmarshal(j, list)
	if err != nil {
		return list, err
	}

	return list, nil
}
