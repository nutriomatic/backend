package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func SendRequest[T any, R any](url string, requestData T) (R, error) {
	var responseData R

	// Marshal the request data into JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return responseData, errors.New("error in marshalling: " + err.Error())
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return responseData, errors.New("error in creating request: " + err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return responseData, errors.New("error in sending request: " + err.Error())
	}
	defer resp.Body.Close()

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		return responseData, errors.New("received non-200 response code: " + http.StatusText(resp.StatusCode))
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return responseData, errors.New("error in reading response: " + err.Error())
	}

	log.Printf("Response body: %s\n", body)

	// Unmarshal the response body into the response data
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return responseData, errors.New("error in unmarshalling: " + err.Error())
	}

	return responseData, nil
}
