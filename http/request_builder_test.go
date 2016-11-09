package http

import "testing"

func Test_NewGetRequestBuilder(t *testing.T) {
	request := NewGetRequestBuilder("http://google.com").Build()

	expectedMethod := "GET"
	expectedURL := "http://google.com"

	if request.Method != expectedMethod {
		t.Errorf("Expected: %v, Actual: %v", expectedMethod, request.Method)
	}

	if request.URL.String() != expectedURL {
		t.Errorf("Expected: %v, Actual: %v", expectedURL, request.URL.String())
	}
}

func Test_NewPostRequestBuilder(t *testing.T) {
	request := NewPostRequestBuilder("http://google.com").Build()

	expectedMethod := "POST"
	expectedURL := "http://google.com"

	if request.Method != "POST" {
		t.Errorf("Expected: %v, Actual: %v", expectedMethod, request.Method)
	}

	if request.URL.String() != "http://google.com" {
		t.Errorf("Expected: %v, Actual: %v", expectedURL, request.URL.String())
	}
}

func Test_WithType(t *testing.T) {
	request := NewRequestBuilder().WithType("POST").Build()

	expectedMethod := "POST"

	if request.Method != expectedMethod {
		t.Errorf("Expected: %v, Actual: %v", expectedMethod, request.Method)
	}
}

func Test_WithURL(t *testing.T) {
	request := NewRequestBuilder().WithURL("http://google.com").Build()

	expectedURL := "http://google.com"

	if request.URL.String() != expectedURL {
		t.Errorf("Expected: %v, Actual: %v", expectedURL, request.URL.String())
	}
}
