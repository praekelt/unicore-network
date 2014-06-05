package main

import (
	"github.com/go-martini/martini"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	request, _ := http.NewRequest("GET", "http://localhost:3000/", nil)
	response := httptest.NewRecorder()
	m := martini.Classic()
	m.Get("/", Get)
	m.ServeHTTP(response, request)

	if response.Code != 200 {
		t.Error("Expected 200 response code, got: %s", response.Code)
	}

	if response.Body.String() != "Hello world!" {
		t.Error("Unexpected response: %s", response.Body.String())
	}
}
