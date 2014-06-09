package server

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

type DB struct {
	Network  string
	Address  string
	Database int32
}

func (db *DB) Handler() martini.Handler {
	// NOTE: On a low level this uses net.Dial, see:
	//		 http://golang.org/pkg/net/#Dial for details on network & addr
	conn, err := redis.DialTimeout(db.Network, db.Address, time.Duration(10)*time.Second)
	if err != nil {
		panic(err)
	}

	// select db
	conn.Cmd("select", db.Database)

	return func(c martini.Context) {
		// make available to subsequent handlers
		c.Map(conn)
		c.Next()
	}
}
