package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type YonomaClient struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

func NewYonomaClient(apiKey string) *YonomaClient {
	return &YonomaClient{
		apiKey:  apiKey,
		baseURL: "https://api.yonoma.io/v1/",
		client:  &http.Client{},
	}
}

func (yc *YonomaClient) Request(method, endpoint string, data interface{}) (map[string]interface{}, error) {
	url := yc.baseURL + endpoint
	var requestBody []byte
	var err error

	if data != nil {
		requestBody, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+yc.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := yc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.New(string(body))
	}

	var responseData map[string]interface{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func main() {
	apiKey := "EY947F5EGSZ5TTPV2SEIGSD7F5OI4VFJFMGZZHZK"
	client := NewYonomaClient(apiKey)

	response, err := client.Request("GET", "tags/list", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", response)
}
