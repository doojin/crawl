package http

import "net/http"

// RequestBuilder constructs HTTP requests
type RequestBuilder struct {
	requestType string
	url         string
}

// NewRequestBuilder returns new instance of request builder
func NewRequestBuilder() *RequestBuilder {
	return new(RequestBuilder)
}

// NewGetRequestBuilder returns request builder for building GET request
func NewGetRequestBuilder(url string) *RequestBuilder {
	return NewRequestBuilder().WithType("GET").WithURL(url)
}

// NewPostRequestBuilder returns request builder for building POST request
func NewPostRequestBuilder(url string) *RequestBuilder {
	return NewRequestBuilder().WithType("POST").WithURL(url)
}

// WithType sets request type
func (rb *RequestBuilder) WithType(requestType string) *RequestBuilder {
	rb.requestType = requestType
	return rb
}

// WithURL sets request URL
func (rb *RequestBuilder) WithURL(url string) *RequestBuilder {
	rb.url = url
	return rb
}

// Build returns HTTP request
func (rb *RequestBuilder) Build() *http.Request {
	request, _ := http.NewRequest(rb.requestType, rb.url, nil)
	return request
}
