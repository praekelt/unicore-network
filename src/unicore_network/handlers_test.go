package unicore_network

import (
	"github.com/go-martini/martini"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	request, _ := http.NewRequest("GET", "/foo", nil)
	response := httptest.NewRecorder()
	m := martini.Classic()
	m.Use(DB())
	m.Get("/:id", Get)
	m.ServeHTTP(response, request)

	if response.Code != 200 {
		t.Errorf("Expected 200 response code, got: %s", response.Code)
	}

	if response.Body.String() != "Hello world! foo" {
		t.Errorf("Unexpected response: %s", response.Body.String())
	}
}
