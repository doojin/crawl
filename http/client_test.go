package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_Execute(t *testing.T) {
	client := NewClient()
	setUpClientTest("google response")

	response := client.Execute(NewGetRequestBuilder("http://google.com"))
	expectedResponse := "google response"

	if response != expectedResponse {
		t.Errorf("Expected: %v, Actual: %v", expectedResponse, response)
	}
}

func setUpClientTest(mockedResponse string) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, mockedResponse)
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	http.DefaultClient.Transport = transport
}
