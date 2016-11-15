package http

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_NewGetRequestBuilder(t *testing.T) {
	request := NewGetRequestBuilder("http://google.com").Build()

	assert.Equal(t, "GET", request.Method)
	assert.Equal(t, "http://google.com", request.URL.String())
}

func Test_NewPostRequestBuilder(t *testing.T) {
	request := NewPostRequestBuilder("http://google.com").Build()

	assert.Equal(t, "POST", request.Method)
	assert.Equal(t, "http://google.com", request.URL.String())
}

func Test_WithType(t *testing.T) {
	request := NewRequestBuilder().WithType("POST").Build()

	assert.Equal(t, "POST", request.Method)
}

func Test_WithURL(t *testing.T) {
	request := NewRequestBuilder().WithURL("http://google.com").Build()

	assert.Equal(t, "http://google.com", request.URL.String())
}
