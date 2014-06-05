package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFoo(t *testing.T) {
	m := New()
	request, _ := http.NewRequest("GET", "/foo", nil)
	response := httptest.NewRecorder()

	m.ServeHTTP(response, request)

	if response.Code != 200 {
		t.Error("Expected 200 response code, got:", response.Code)
	}

	if response.Body.String() != "Hello world! foo" {
		t.Error("Unexpected response:", response.Body.String())
	}
}
