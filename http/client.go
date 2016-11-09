package http

import (
	"io/ioutil"
	"net/http"
)

var defaultClient *Client

func init() {
	defaultClient = new(Client)
}

// Client executes HTTP requests
type Client struct {
}

// NewClient returns new instance of Client
func NewClient() Client {
	return Client{}
}

// DefaultClient returns default instance of Client
func DefaultClient() *Client {
	return defaultClient
}

// Execute executes HTTP request built by request builder
func (c *Client) Execute(rb *RequestBuilder) string {
	request := rb.Build()
	resp, _ := http.DefaultClient.Do(request)
	return c.readResponse(resp)
}

func (c *Client) readResponse(resp *http.Response) string {
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
