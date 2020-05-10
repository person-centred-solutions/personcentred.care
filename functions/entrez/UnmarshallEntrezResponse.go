
import (
	"encoding/json"
	"encoding/xml"
)

func UnmarshallEntrezResponse(response string) (string, error) {
	var data Data
	err := xml.Unmarshal([]byte(response), &data)
	if err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return jsonData
}
