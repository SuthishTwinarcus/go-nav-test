package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type GroupsYonomaClient struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

func NewGroupYonomaClient(apiKey string) *GroupsYonomaClient {
	return &GroupsYonomaClient{
		apiKey:  apiKey,
		baseURL: "http://localhost:8080/v1/",
		client:  &http.Client{},
	}
}

func (yc *GroupsYonomaClient) Request(method, endpoint string, data interface{}) (map[string]interface{}, error) {
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

type Lists struct {
	client *GroupsYonomaClient
}

// NewLists creates a new Lists instance.
func NewLists(client *GroupsYonomaClient) *Lists {
	return &Lists{client: client}
}

// Create a new list
func (g *Lists) Create(listData map[string]interface{}) (map[string]interface{}, error) {
	endpoint := "lists/create"
	return g.client.Request("POST", endpoint, listData)
}

// List all Lists
func (g *Lists) List() (map[string]interface{}, error) {
	endpoint := "lists/list"
	return g.client.Request("GET", endpoint, nil)
}

// Retrieve a specific list by ID
func (g *Lists) Retrieve(listID string) (map[string]interface{}, error) {
	endpoint := "lists/" + listID
	return g.client.Request("GET", endpoint, nil)
}

// Update a list
func (g *Lists) Update(listID string, listData map[string]interface{}) (map[string]interface{}, error) {
	endpoint := "lists/" + listID + "/update"
	return g.client.Request("POST", endpoint, listData)
}

// Delete a list
func (g *Lists) Delete(listID string) (map[string]interface{}, error) {
	endpoint := "lists/" + listID + "/delete"
	return g.client.Request("POST", endpoint, nil)
}

// func main() {
// 	apiKey := "EY947F5EGSZ5TTPV2SEIGSD7F5OI4VFJFMGZZHZK"
// 	client := NewGroupYonomaClient(apiKey)

// 	// Initialize Lists Service
// 	lists := NewLists(client)

// 	// Define the list data
// 	listData := map[string]interface{}{
// 		"list_name": "naveen chand",
// 	}

// 	// Call Create method
// 	response, err := lists.Create(listData)
// 	if err != nil {
// 		log.Fatalf("Error creating list: %v", err)
// 	}

// 	// Print success response
// 	fmt.Println("List Created Successfully:", response)
// }

// func main() {
// 	// Initialize the YonomaClient with your API key
// 	client := NewYonomaClient("EY947F5EGSZ5TTPV2SEIGSD7F5OI4VFJFMGZZHZK")

// 	// Create a Lists instance
// 	Lists := NewLists(client)

// 	// Call the List method to fetch all Lists
// 	listList, err := Lists.List()
// 	if err != nil {
// 		log.Fatalf("Error fetching Lists: %v", err)
// 	}

// 	// Print the list of Lists
// 	fmt.Println("Lists List:", listList)
// }

// func main() {
// 	// Initialize the API Client
// 	apiKey := "EY947F5EGSZ5TTPV2SEIGSD7F5OI4VFJFMGZZHZK"
// 	client := NewYonomaClient(apiKey)

// 	// Initialize Lists Service
// 	lists := NewLists(client)

// 	// List ID to retrieve
// 	listID := "WAFZGSN76D" // Replace with actual list ID

// 	// Call Retrieve method
// 	response, err := lists.Retrieve(listID)
// 	if err != nil {
// 		log.Fatalf("Error retrieving list: %v", err)
// 	}

// 	// Print retrieved list details
// 	fmt.Println("List Retrieved Successfully:", response)
// }

// func main() {
// 	// Initialize the API Client
// 	apiKey := "EY947F5EGSZ5TTPV2SEIGSD7F5OI4VFJFMGZZHZK"
// 	client := NewYonomaClient(apiKey)

// 	// Initialize Lists Service
// 	lists := NewLists(client)

// 	// List ID to update
// 	listID := "WAFZGSN76D" // Replace with actual list ID

// 	// Define the updated list data
// 	updateData := map[string]interface{}{
// 		"list_name": "naveen list",
// 	}

// 	// Call Update method
// 	response, err := lists.Update(listID, updateData)
// 	if err != nil {
// 		log.Fatalf("Error updating list: %v", err)
// 	}

// 	// Print success response
// 	fmt.Println("List Updated Successfully:", response)
// }

// func main() {
// 	// Initialize the API Client
// 	apiKey := "EY947F5EGSZ5TTPV2SEIGSD7F5OI4VFJFMGZZHZK"
// 	client := NewYonomaClient(apiKey)

// 	// Initialize Lists Service
// 	lists := NewLists(client)

// 	// List ID to delete
// 	listID := "WAFZGSN76D" // Replace with actual list ID

// 	// Call Delete method
// 	response, err := lists.Delete(listID)
// 	if err != nil {
// 		log.Fatalf("Error deleting list: %v", err)
// 	}

// 	// Print success response
// 	fmt.Println("List Deleted Successfully:", response)
// }
