package server

import (
	"github.com/fzzy/radix/redis"
	"time"
)

type DB struct {
	Network    string
	Address    string
	Database   int
	Connection redis.Client
}

func (db *DB) Connect() (*redis.Client, error) {
	// NOTE: On a low level this uses net.Dial, see:
	//		 http://golang.org/pkg/net/#Dial for details on network & addr
	conn, err := redis.DialTimeout(db.Network, db.Address, time.Duration(10)*time.Second)
	if err != nil {
		return conn, err
	}
	result := conn.Cmd("select", db.Database)
	if result.Err != nil {
		return conn, err
	}
	return conn, err
}
