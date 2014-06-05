package server

import (
	"github.com/go-martini/martini"
)

func New() *martini.ClassicMartini {
	m := martini.Classic()
	m.Use(DB())
	m.Get("/identity", GetIdentity)
	return m
}
