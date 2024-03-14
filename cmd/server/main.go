package main

import (
	"log"

	"github.com/mahonyodhran/proglog/internal/server"
)

func main() {
	srv := server.newHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
