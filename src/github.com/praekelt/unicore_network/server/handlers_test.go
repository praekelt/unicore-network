package server

import (
	"encoding/json"
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

	ident := Ident{}
	err := json.Unmarshal(response.Body.Bytes(), &ident)
	if err != nil {
		t.Error(err)
	}
	if ident != CreateIdentity("identity", "localhost", "test node") {
		t.Error("Unexpected identity returned", ident)
	}
}
