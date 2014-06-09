package main

import (
	"flag"
	"github.com/praekelt/unicore_network/server"
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

	identity, _ := server.GetOrCreateIdentity(identity_file)

	m := server.New(identity)
	log.Fatal(http.ListenAndServe(address, m))
}
