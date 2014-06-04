package main

import (
	"github.com/go-martini/martini"
)

func New() *martini.ClassicMartini {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
	return m
}

func main() {
	New()
}
