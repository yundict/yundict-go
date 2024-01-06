package yundict

import (
	"io"
	"net/http"
)

const (
	apiTokenHeader = "X-Api-Token"
	defaultBaseURL = "https://api.yundict.com"
)

type Client struct {
	BaseURL string
	Token   string
	Keys    *KeysService
}

func NewClient(token string) *Client {
	c := &Client{
		BaseURL: defaultBaseURL,
		Token:   token,
	}

	c.Keys = &KeysService{client: c}

	return c
}

type service struct {
	client *Client
}

// Get makes a GET request to the given path
func (c *Client) Get(path string) ([]byte, error) {

	url := c.BaseURL + path
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
