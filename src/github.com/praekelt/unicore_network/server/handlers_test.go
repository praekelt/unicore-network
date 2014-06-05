package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func do_request(request *http.Request) *httptest.ResponseRecorder {
	ident := CreateIdentity("identity", "localhost", "test node")
	m := New(ident)
	response := httptest.NewRecorder()
	m.ServeHTTP(response, request)
	return response
}

func TestGetIdentity(t *testing.T) {

	request, _ := http.NewRequest("GET", "/identity", nil)
	response := do_request(request)

	if response.Code != 200 {
		t.Error("Expected 200 response code, got:", response.Code)
	}

	if response.Body.String() != "Hello world!" {
		t.Error("Unexpected response:", response.Body.String())
	}
}
