package server

import (
	"encoding/json"
	"github.com/fzzy/radix/redis"
	"net/http"
)

func GetIdentity(request *http.Request, db *redis.Client, identity Ident) string {
	bytes, _ := json.Marshal(identity)
	return string(bytes)
}
