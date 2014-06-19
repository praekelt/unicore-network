package server

import (
	"fmt"
	"github.com/fzzy/radix/redis"
	"github.com/go-martini/martini"
	"net/http"
)

func (s *Server) GetOwnIdentity(response http.ResponseWriter, request *http.Request) (int, string) {
	json, _ := s.Identity.ToString()
	return http.StatusOK, json
}

func (s *Server) PutNodeIdentity(response http.ResponseWriter, request *http.Request, db *redis.Client, params martini.Params) (int, string) {
	ident, err := NewIdentFromReader(request.Body)
	if err != nil {
		panic(err)
	}

	// if the signature param in the URL doesn't match the signature in the JSON
	// payload then we should raise an error.
	if params["signature"] != ident.Signature {
		http.Error(response, "Signature in payload must match URL", http.StatusBadRequest)
	}

	_, put_err := s.PutIdent(db, ident)
	if put_err != nil {
		panic(put_err)
	}

	// NOTE:    As far as I can tell Martini doesn't have the concept of named
	//          URLs like Django has which forces me to duplicate URL path names
	//          here and in server.go
	response.Header().Set("Location", fmt.Sprintf("/network/%s", ident.Signature))
	return http.StatusCreated, ""
}

func (s *Server) GetNodeIdentity(response http.ResponseWriter, request *http.Request, db *redis.Client, params martini.Params) (int, string) {
	ident, err := s.GetIdent(db, params["signature"])
	if err != nil {
		panic(err)
	}
	json_string, _ := ident.ToString()
	return http.StatusOK, json_string
}
