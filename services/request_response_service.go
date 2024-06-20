package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func SendRequest[T any, R any](url string, requestData T) (R, error) {
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return *new(R), err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return *new(R), err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return *new(R), err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return *new(R), err
	}

	var responseData R
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return *new(R), err
	}

	return responseData, nil
}
