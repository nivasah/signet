package main

import (
	"log"

	"github.com/nivasah/signet/internal/server"
)

func main() {
	cfg := server.Config{
		Port: "8080",
	}
	log.Fatal(server.Run(&cfg))
}
