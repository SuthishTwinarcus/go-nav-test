package models

import "net/http"

// User struct defined in models/user.go
type YonomaClient struct {
	APIKey  string
	BaseURL string
	Client  *http.Client
}

type Contacts struct {
	Client *YonomaClient
}
