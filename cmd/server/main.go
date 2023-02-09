package main

import (
	"log"

	"github.com/brodiep21/distserv/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
