package entrez

import (
	"encoding/json"
	"encoding/xml"
)

func UnmarshallEntrezResponse(response string) (string, error) {
	var data EntrezResponse
	err := xml.Unmarshal([]byte(response), &data)
	if err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
