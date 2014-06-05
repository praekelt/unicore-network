package main

import (
	"github.com/go-martini/martini"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestHandlers(c *C) {
	request, _ := http.NewRequest("GET", "http://localhost:3000/", nil)
	response := httptest.NewRecorder()
	m := martini.Classic()
	m.Get("/", Get)
	m.ServeHTTP(response, request)

	c.Assert(response.Code, Equals, 200)
	c.Assert(response.Body.String(), Equals, "Hello world!")
}
