package server

import (
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
	m.Get("/network", s.GetNodeIdentityIndex)
	m.Put("/network/:signature", s.PutNodeIdentity)
	m.Get("/network/:signature", s.GetNodeIdentity)
	m.Delete("/network/:signature", s.DeleteNodeIdentity)
	return m
}

func key(prefix string, parts ...string) string {
	return strings.Join(append([]string{prefix}, parts...), "/")
}

func node_key(parts ...string) string {
	return key("node", parts...)
}

func (s *Server) PutIdent(db *redis.Client, identity Ident) (Ident, error) {
	data, err := identity.ToJson()
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

func (s *Server) GetIdentIndex(db *redis.Client, start int, stop int) ([]Ident, error) {
	zrange, err := db.Cmd("zrange", "nodes", start, stop).List()
	if err != nil {
		return []Ident{}, err
	}

	idents := make([]Ident, len(zrange))
	for index, signature := range zrange {
		ident, err := s.GetIdent(db, signature)
		if err != nil {
			return []Ident{}, err
		}
		idents[index] = ident
	}
	return idents, nil
}

func (s *Server) DeleteIdent(db *redis.Client, signature string) (Ident, error) {
	ident := Ident{}
	get := db.Cmd("get", node_key(signature))
	if get.Err != nil {
		return ident, get.Err
	}

	if get.Type == redis.NilReply {
		return ident, errors.New(fmt.Sprintf("Ident %s does not exist.", signature))
	}

	str, err := get.Str()
	if err != nil {
		panic(err)
	}
	del := db.Cmd("del", node_key(signature))
	if del.Err != nil {
		return ident, del.Err
	}

	if num_deleted, _ := del.Int(); num_deleted == 0 {
		return ident, errors.New(fmt.Sprintf("Deleting %s failed.", signature))
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
