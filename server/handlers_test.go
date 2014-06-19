package server

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func new_server() Server {
	ident := CreateIdentity("identity", "localhost", "test node")
	db := &DB{Network: "tcp", Address: "127.0.0.1:6379", Database: 0}
	server := Server{Identity: ident, Db: db}
	return server
}

func do_request(request *http.Request) *httptest.ResponseRecorder {
	server := new_server()
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

	ident, err := NewIdentFromReader(response.Body)
	if err != nil {
		t.Error(err)
	}
	if ident != CreateIdentity("identity", "localhost", "test node") {
		t.Error("Unexpected identity returned", ident)
	}
}

func TestPutNodeIdentity(t *testing.T) {
	ident := CreateIdentity("foo", "bar", "baz")
	b, _ := ident.ToJson()
	request, _ := http.NewRequest("PUT", "/network/foo", bytes.NewReader(b))
	response := do_request(request)
	if response.Code != http.StatusCreated {
		t.Error("Unexpected HTTP response", response.Code)
	}
	location_header := response.Header().Get("Location")
	if location_header != "/network/foo" {
		t.Error("Unexpected Location header", location_header)
	}
}

func TestGetNodeIdentity(t *testing.T) {
	ident := CreateIdentity("foo", "bar", "baz")
	server := new_server()
	conn, _ := server.Db.Connect()
	server.PutIdent(conn, ident)
	request, _ := http.NewRequest("GET", "/network/foo", nil)
	response := do_request(request)
	if response.Code != http.StatusOK {
		t.Error("Unexpected HTTP response", response.Code)
	}
	found_ident, _ := NewIdentFromReader(response.Body)
	if found_ident != ident {
		t.Error("Unexpected Ident returned", found_ident)
	}
	fmt.Println(found_ident)
}
