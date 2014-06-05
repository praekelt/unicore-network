package unicore_network

import (
	"fmt"
	"github.com/fzzy/radix/redis"
	"github.com/go-martini/martini"
	"net/http"
)

func Get(request *http.Request, db *redis.Client, parms martini.Params) string {
	fmt.Println("redis", db)
	return fmt.Sprintf("Hello world! %s", parms["id"])
}
