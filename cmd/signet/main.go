package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nivasah/signet/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("SIGNET_PORT")
	cfg := server.Config{
		Port: port,
	}
	log.Fatal(server.Run(&cfg))
}
