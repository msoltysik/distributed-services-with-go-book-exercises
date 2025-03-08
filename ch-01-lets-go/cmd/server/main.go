package main

import (
	"log"

	"github.com/msoltysik/distributed-services-with-go-book-exercises/ch-01-lets-go/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
