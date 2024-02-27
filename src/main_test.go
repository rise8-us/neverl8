package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"


	"github.com/stretchr/testify/assert"
	"github.com/drewfugate/neverl8"
)

func TestHelloWorldEndpoint(t *testing.T) {
	// Create a new request to the endpoint (assuming it's "/")
	req, err := http.NewRequest("GET", "/helloworld", nil)
	assert.NoError(t, err, "should not error creating request")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler using your application's router
	handler := http.HandlerFunc(main.HelloWorldHandler)

	// Serve the HTTP request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code, "status code should be OK")

	// Check the response body
	expected := "Hello, World!"
	assert.Equal(t, expected, rr.Body.String(), "response body should match expected")
}
