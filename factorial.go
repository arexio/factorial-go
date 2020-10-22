package factorial

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
)

const factorialAPI = "https://api.factorialhr.com"

// New builds a Factorial client from the provided accessToken and options.
func New(opts ...Option) (*Client, error) {
	c := &Client{
		apiURL: factorialAPI,
	}
	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

// WithOAuth2Client provides a custom http client to the client.
func WithOAuth2Client(cli *http.Client) func(*Client) {
	return func(c *Client) {
		c.Client = cli
	}
}

// WithAPIURL sets the API URL for the client.
// Useful for testing.
func WithAPIURL(url string) func(*Client) {
	return func(c *Client) {
		c.apiURL = url
	}
}

// Option defines an option for a Client.
type Option func(*Client)

// Client for the Factorial API.
type Client struct {
	*http.Client
	apiURL string
}

func (c Client) delete(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, c.apiURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return c.Do(req)
}

func (c Client) get(endpoint string, q url.Values) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.apiURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	if q != nil {
		req.URL.RawQuery = q.Encode()
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return c.Do(req)
}

func (c Client) post(endpoint string, body []byte) (*http.Response, error) {
	log.Println("[DEBUG] url", c.apiURL+endpoint, string(body))
	req, err := http.NewRequest(http.MethodPost, c.apiURL+endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return c.Do(req)
}

func (c Client) put(endpoint string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, c.apiURL+endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return c.Do(req)
}
