package main

import (
	"flag"
	"github.com/go-martini/martini"
	"github.com/praekelt/unicore_network/server"
	"log"
	"net/http"
)

var m *martini.ClassicMartini

func init() {
	m = server.New()
}

func main() {
	address := flag.String("address", ":8080", "The address to listen on.")
	flag.Parse()

	log.Printf("Listening on %s", *address)
	err := http.ListenAndServe(*address, m)
	if err != nil {
		log.Fatal(err)
	}
}
