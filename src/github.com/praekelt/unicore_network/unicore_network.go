package main

import (
	"flag"
	"github.com/praekelt/unicore_network/server"
	"log"
	"net/http"
)

func main() {
	address := flag.String("address", ":8080", "The address to listen on.")
	identity_file := flag.String("identity", "~/.uc_identity.yaml", "Which identity file to use.")
	flag.Parse()

	log.Printf("Listening on %s", *address)

	identity, _ := server.GetOrCreateIdentity(*identity_file)

	m := server.New(identity)
	err := http.ListenAndServe(*address, m)
	if err != nil {
		log.Fatal(err)
	}
}
