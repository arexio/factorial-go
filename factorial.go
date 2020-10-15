package factorial

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

const factorialAPI = "https://api.factorialhr.com"

// New builds a Factorial client from the provided accessToken and options.
func New(accessToken string, opts ...Option) (*Client, error) {
	c := &Client{
		apiURL: factorialAPI,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		accessToken: accessToken,
	}
	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

// OptionHTTPClient provides a custom http client to the client.
func OptionHTTPClient(cli httpClient) func(*Client) {
	return func(c *Client) {
		c.httpClient = cli
	}
}

// OptionAPIURL sets the API URL for the client.
// Useful for testing.
func OptionAPIURL(url string) func(*Client) {
	return func(c *Client) {
		c.apiURL = url
	}
}

// Option defines an option for a Client.
type Option func(*Client)

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Client for the Factorial API.
type Client struct {
	accessToken string
	apiURL      string
	httpClient  httpClient
}

func (c Client) do(req *http.Request) (*http.Response, error) {
	authHeader := fmt.Sprintf("Bearer %s", c.accessToken)

	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error doing request: %s", resp.Status)
	}

	return resp, nil
}

func (c Client) get(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.apiURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

func (c Client) post(endpoint string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, c.apiURL+endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

func (c Client) put(endpoint string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, c.apiURL+endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	return c.do(req)
}
