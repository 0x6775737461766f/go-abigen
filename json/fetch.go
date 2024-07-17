package json

import (
	"encoding/json"
	"io"
	"net/http"
)

// JSON structure from the API
type Data struct {
	Data  string `json:"data"`
	Valor string `json:"valor"`
}

// JSON data from a URL and returns an array of Data
func FetchData(url string) ([]Data, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []Data
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
