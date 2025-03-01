package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ContactsYonomaClient struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

func NewContactsYonomaClient(apiKey string) *ContactsYonomaClient {
	return &ContactsYonomaClient{
		apiKey:  apiKey,
		baseURL: "http://localhost:8080/v1/",
		client:  &http.Client{},
	}
}

func (yc *ContactsYonomaClient) Request(method, endpoint string, data interface{}) (map[string]interface{}, error) {
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

type Contacts struct {
	client *ContactsYonomaClient
}

func NewContacts(client *ContactsYonomaClient) *Contacts {
	return &Contacts{client: client}
}

func (c *Contacts) Create(groupId string, contactData map[string]interface{}) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("contacts/%s/create", groupId)
	return c.client.Request("POST", endpoint, contactData)
}

func (c *Contacts) Update(groupId, contactId string, contactData map[string]interface{}) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("contacts/%s/status/%s", groupId, contactId)
	return c.client.Request("POST", endpoint, contactData)
}

func (c *Contacts) AddTag(contactId string, contactData map[string]interface{}) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("contacts/tags/%s/add", contactId)
	return c.client.Request("POST", endpoint, contactData)
}

func (c *Contacts) RemoveTag(contactId string, contactData map[string]interface{}) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("contacts/tags/%s/remove", contactId)
	return c.client.Request("POST", endpoint, contactData)
}

// func main() {
// 	apiKey := "EY947F5EGSZ5TTPV2SEIGSD7F5OI4VFJFMGZZHZK"
// 	client := NewContactsYonomaClient(apiKey)
// 	contacts := NewContacts(client)

// 	contactData := map[string]interface{}{
// 		"status": "Subscribed",
// 		"email":  "naveenac12@example.com",
// 	}

// 	response, err := contacts.Create("WAFZGSN76D", contactData)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println("Response:", response)

// 	// Update Contact Example
// 	// updateData := map[string]interface{}{
// 	// 	"status": "Unsubscribed",
// 	// }
// 	// updateResponse, err := contacts.Update("27OE7G54R5", "MSHX71HTPU", updateData)
// 	// if err != nil {
// 	// 	fmt.Println("Error updating contact:", err)
// 	// 	return
// 	// }
// 	// fmt.Println("Update Response:", updateResponse)

// 	// Add Tag to Contact Example
// 	// tagData := map[string]interface{}{
// 	// 	"tag_id": "CBJJJXOGGO",
// 	// }
// 	// tagResponse, err := contacts.AddTag("MSHX71HTPU", tagData)
// 	// if err != nil {
// 	// 	fmt.Println("Error adding tag to contact:", err)
// 	// 	return
// 	// }
// 	// fmt.Println("Add Tag Response:", tagResponse)

// 	// Remove Tag from Contact Example
// 	// removeTagData := map[string]interface{}{
// 	// 	"tag_id": "CBJJJXOGGO",
// 	// }
// 	// removeTagResponse, err := contacts.RemoveTag("MSHX71HTPU", removeTagData)
// 	// if err != nil {
// 	// 	fmt.Println("Error removing tag from contact:", err)
// 	// 	return
// 	// }
// 	// fmt.Println("Remove Tag Response:", removeTagResponse)
// }
