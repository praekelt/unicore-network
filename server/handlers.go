package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) GetOwnIdentity(response http.ResponseWriter, request *http.Request) (int, string) {
	bytes, _ := json.Marshal(s.Identity)
	return http.StatusOK, string(bytes)
}

func (s *Server) PutNodeIdentity(response http.ResponseWriter, request *http.Request) (int, string) {
	ident := Ident{}
	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&ident)

	conn, err := s.Db.Connect()
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	err = s.Db.Save(conn, ident)
	if err != nil {
		panic(err)
	}

	// NOTE:    As far as I can tell Martini doesn't have the concept of named
	//          URLs like Django has which forces me to duplicate URL path names
	//          here and in server.go
	response.Header().Set("Location", fmt.Sprintf("/network/%s", ident.Signature))
	return http.StatusCreated, ""
}
