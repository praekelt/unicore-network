package server

import (
	"encoding/json"
	"github.com/fzzy/radix/redis"
	"time"
)

type DB struct {
	Network  string
	Address  string
	Database int
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

func (db *DB) Save(conn *redis.Client, ident Ident) error {
	data, err := json.Marshal(ident)
	if err != nil {
		panic(err)
	}
	result := conn.Cmd("zadd", "nodes", time.Now().Unix(), data)
	return result.Err
}
