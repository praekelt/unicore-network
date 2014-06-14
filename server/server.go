package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fzzy/radix/redis"
	"github.com/go-martini/martini"
	"strings"
	"time"
)

type Server struct {
	Identity Ident
	Db       *DB
}

func (s *Server) New() *martini.ClassicMartini {
	m := martini.Classic()
	m.Use(s.Handler())
	m.Get("/identity", s.GetOwnIdentity)
	m.Put("/network/:signature", s.PutNodeIdentity)
	m.Get("/network/:signature", s.GetNodeIdentity)
	return m
}

func key(prefix string, parts ...string) string {
	return strings.Join(append([]string{prefix}, parts...), "/")
}

func node_key(parts ...string) string {
	return key("node", parts...)
}

func (s *Server) PutIdent(db *redis.Client, identity Ident) (Ident, error) {
	data, err := json.Marshal(identity)
	if err != nil {
		return identity, err
	}
	zadd := db.Cmd("zadd", "nodes", time.Now().Unix(), identity.Signature)
	if zadd.Err != nil {
		return identity, zadd.Err
	}

	set := db.Cmd("set", node_key(identity.Signature), data)
	if set.Err != nil {
		return identity, set.Err
	}
	return identity, nil
}

func (s *Server) GetIdent(db *redis.Client, signature string) (Ident, error) {
	ident := Ident{}
	get := db.Cmd("get", node_key(signature))
	if get.Err != nil {
		fmt.Println("1")
		return ident, get.Err
	}

	if get.Type == redis.NilReply {
		return ident, errors.New(fmt.Sprintf("Ident %s does not exist.", signature))
	}

	str, err := get.Str()
	if err != nil {
		panic(err)
	}
	return NewIdentFromString(str)
}

func (s *Server) Handler() martini.Handler {
	return func(c martini.Context) {
		conn, err := s.Db.Connect()
		if err != nil {
			panic(err)
		}
		c.Map(conn)
		defer conn.Close()
		c.Next()
	}
}
