package main

// YonomaClient represents the API client
// type YonomaClient struct {
// 	apiKey string
// }

// NewYonomaClient initializes a new YonomaClient
// func NewYonomaClient(apiKey string) *YonomaClient {
// 	return &YonomaClient{apiKey: apiKey}
// }

// Contacts represents contact management
// type Contacts struct {
// 	client *YonomaClient
// }

// Lists represents group management
// type Lists struct {
// 	client *YonomaClient
// }

// // Tags represents tag management
// type Tags struct {
// 	client *YonomaClient
// }

// ApiClient is the main struct that holds all components
type ApiClient struct {
	client   *YonomaClient
	Contacts *Contacts
	Lists    *Lists
	Tags     *Tags
}

// NewApiClient initializes a new ApiClient
func NewApiClient(apiKey string) *ApiClient {
	client := NewYonomaClient(apiKey)
	contactClient := NewContactsYonomaClient(apiKey)
	listClient := NewGroupYonomaClient(apiKey)
	tagsClient := NewTagsYonomaClient(apiKey)

	return &ApiClient{
		client:   client,
		Contacts: &Contacts{client: contactClient},
		Lists:    &Lists{client: listClient},
		Tags:     &Tags{client: tagsClient},
	}
}

// Example usage
// func main() {
// 	apiKey := "EY947F5EGSZ5TTPV2SEIGSD7F5OI4VFJFMGZZHZK"
// 	apiClient := NewApiClient(apiKey)

// 	fmt.Println("Yonoma API Client initialized")
// 	fmt.Printf("Client: %+v\n", apiClient)
// }
