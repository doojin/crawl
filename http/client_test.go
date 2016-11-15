package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Execute(t *testing.T) {
	client := NewClient()
	setUpClientTest("google response")

	response := client.Execute(NewGetRequestBuilder("http://google.com"))

	assert.Equal(t, "google response", response)
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
