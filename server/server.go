package server

import (
	"github.com/go-martini/martini"
)

type Server struct {
	Identity Ident
	Db       *DB
}

func (s *Server) New() *martini.ClassicMartini {
	m := martini.Classic()
	m.Get("/identity", s.GetOwnIdentity)
	m.Put("/network/:string", s.PutNodeIdentity)
	return m
}
