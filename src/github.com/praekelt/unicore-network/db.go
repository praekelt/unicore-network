package main

import (
	"github.com/fzzy/radix/redis"
	"github.com/go-martini/martini"
	"time"
)

/*
NOTE: http://blog.gopheracademy.com/day-11-martini explains how this stuff
      mostly works. It's a bit like Django middleware but allows one to
      insert extra parameters that are passed along to the handler (view)
*/
func DB() martini.Handler {
	// connect to the db
	conn, err := redis.DialTimeout(
		"tcp", "127.0.0.1:6379",
		time.Duration(10)*time.Second)
	if err != nil {
		panic(err)
	}
	// close the connection when done
	defer conn.Close()

	// select db
	conn.Cmd("select", 0)

	return func(c martini.Context) {
		// make available to subsequent handlers
		c.Map(conn)
		c.Next()
	}
}
