package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func do_request(request *http.Request) *httptest.ResponseRecorder {
	ident := CreateIdentity("identity", "localhost", "test node")
	db := DB{Network: "tcp", Address: "127.0.0.1:6379", Database: 0}
	server := Server{Identity: ident, Db: db}
	martini := server.New()
	response := httptest.NewRecorder()
	martini.ServeHTTP(response, request)
	return response
}

func TestGetOwnIdentity(t *testing.T) {

	request, _ := http.NewRequest("GET", "/identity", nil)
	response := do_request(request)

	if response.Code != http.StatusOK {
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

func TestPutNodeIdentity(t *testing.T) {
	ident := CreateIdentity("foo", "bar", "baz")
	b, _ := json.Marshal(ident)
	request, _ := http.NewRequest("PUT", "/network/foo", bytes.NewReader(b))
	response := do_request(request)
	if response.Code != http.StatusCreated {
		t.Error("Unexpected HTTP response", response)
	}
	location_header := response.Header().Get("Location")
	if location_header != "/network/foo" {
		t.Error("Unexpected Location header", location_header)
	}
}
