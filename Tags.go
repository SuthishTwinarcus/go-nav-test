package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TagsYonomaClient struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

func NewTagsYonomaClient(apiKey string) *TagsYonomaClient {
	return &TagsYonomaClient{
		apiKey:  apiKey,
		baseURL: "http://localhost:8080/v1/",
		client:  &http.Client{},
	}
}

func (yc *TagsYonomaClient) Request(method, endpoint string, data interface{}) (map[string]interface{}, error) {
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

// Tags represents the Yonoma Tags API
type Tags struct {
	client *TagsYonomaClient
}

// NewTags initializes a new Tags instance
func NewTags(client *TagsYonomaClient) *Tags {
	return &Tags{client: client}
}

// Create a new tag
func (t *Tags) Create(tagData map[string]interface{}) (map[string]interface{}, error) {
	return t.client.Request("POST", "tags/create", tagData)
}

// List all tags
func (t *Tags) List() (map[string]interface{}, error) {
	return t.client.Request("GET", "tags/list", nil)
}

// Retrieve details of a specific tag
func (t *Tags) Retrieve(tagID string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("tags/%s", tagID)
	return t.client.Request("GET", endpoint, nil)
}

// Update an existing tag
func (t *Tags) Update(tagID string, tagData map[string]interface{}) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("tags/%s/update", tagID)
	return t.client.Request("POST", endpoint, tagData)
}

// Delete a tag
func (t *Tags) Delete(tagID string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("tags/%s/delete", tagID)
	return t.client.Request("POST", endpoint, nil)
}

// func main() {
// 	client := NewYonomaClient("EY947F5EGSZ5TTPV2SEIGSD7F5OI4VFJFMGZZHZK")
// 	tags := NewTags(client)

// 	// tagData := map[string]interface{}{
// 	// 	"tag_name": "tag12",
// 	// }

// 	// response, err := tags.Create(tagData)
// 	// if err != nil {
// 	// 	log.Fatalf("Error creating tag: %v", err)
// 	// }

// 	// fmt.Println("Tag Created Successfully:", response)

// 	//////tag list/////
// 	// response, err := tags.List()
// 	// if err != nil {
// 	// 	log.Fatalf("Error retrieving tags: %v", err)
// 	// }
// 	// fmt.Println("Tags:", response)

// 	//retrieve tag///////
// 	// tagID := "NVQDIZVR1O" // Replace with actual tag ID
// 	// response, err := tags.Retrieve(tagID)
// 	// if err != nil {
// 	// 	log.Fatalf("Error retrieving tag: %v", err)
// 	// }
// 	// fmt.Println("Tag Details:", response)

// 	///update tag ////////////
// 	// tagID := "NVQDIZVR1O" // Replace with actual tag ID
// 	// updateData := map[string]interface{}{
// 	// 	"tag_name": "naveen tg",
// 	// }

// 	// response, err := tags.Update(tagID, updateData)
// 	// if err != nil {
// 	// 	log.Fatalf("Error updating tag: %v", err)
// 	// }

// 	// fmt.Println("Tag Updated Successfully:", response)

// 	///delete tag /////////
// 	// tagID := "OBOFEK6MPO" // Replace with actual tag ID
// 	// response, err := tags.Delete(tagID)
// 	// if err != nil {
// 	// 	log.Fatalf("Error deleting tag: %v", err)
// 	// }

// 	// fmt.Println("Tag Deleted Successfully:", response)
// }
