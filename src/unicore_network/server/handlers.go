package server

import (
	"fmt"
	"github.com/fzzy/radix/redis"
	"net/http"
)

func GetIdentity(request *http.Request, db *redis.Client) string {
	return fmt.Sprintf("Hello world!")
}
