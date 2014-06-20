package main

import (
	"flag"
	. "github.com/praekelt/unicore-network/server"
	"log"
	"net/http"
	"os/user"
	"path/filepath"
)

func main() {
	var identity_file string
	var address string

	flag.StringVar(&address, "address", ":8080", "The address to listen on.")
	flag.StringVar(&identity_file, "identity", "", "Which identity file to use.")
	flag.Parse()

	log.Printf("Listening on %s", address)

	if identity_file == "" {
		user, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		identity_file = filepath.Join(user.HomeDir, ".uc_identity.yaml")
	}

	identity, _ := GetOrCreateIdentity(identity_file)

	// TODO: make this configurable again
	db := &DB{Network: "tcp", Address: "127.0.0.1:6379", Database: 0}
	server := Server{Identity: identity, Db: db}

	log.Fatal(http.ListenAndServe(address, server.New()))
}
