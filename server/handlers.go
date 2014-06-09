package server

import (
	"encoding/json"
	"fmt"
	"github.com/fzzy/radix/redis"
	"io/ioutil"
	"net/http"
	"time"
)

func (s *Server) GetOwnIdentity(response http.ResponseWriter, request *http.Request) (int, string) {
	bytes, _ := json.Marshal(s.Identity)
	return http.StatusOK, string(bytes)
}

func (s *Server) PutNodeIdentity(response http.ResponseWriter, request *http.Request, db *redis.Client) (int, string) {
	request_bytes, _ := ioutil.ReadAll(request.Body)
	request_body := string(request_bytes)

	ident := Ident{}
	json.Unmarshal(request_bytes, &ident)

	result := db.Cmd("zadd", time.Now().Unix(), request_body)
	i, _ := result.Int64()
	fmt.Println("result", i)

	// NOTE:    As far as I can tell Martini doesn't have the concept of named
	//          URLs like Django has which forces me to duplicate URL path names
	//          here and in server.go
	response.Header().Set("Location", fmt.Sprintf("/network/%s", ident.Signature))
	return http.StatusCreated, ""
}
