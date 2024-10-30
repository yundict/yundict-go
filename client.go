package yundict

import (
	"io"
	"net/http"
)

const (
	// apiTokenHeader is the header key for the API token
	apiTokenHeader = "X-Api-Token"

	// defaultEndpoint is the default API endpoint
	defaultEndpoint = "https://api.yundict.com"
)

// Client is the client for the Yundict API
type Client struct {
	Endpoint string
	Token    string
	Keys     *KeysService
}

// NewClient creates a new Yundict API client
func NewClient(token string) *Client {
	c := &Client{
		Endpoint: defaultEndpoint,
		Token:    token,
	}

	c.Keys = &KeysService{client: c}

	return c
}

// KeysService is the service for the keys endpoint
type service struct {
	client *Client
}

// Get makes a GET request to the given path
func (c *Client) Get(path string) ([]byte, error) {
	url := c.Endpoint + path
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add(apiTokenHeader, c.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body, nil
}
