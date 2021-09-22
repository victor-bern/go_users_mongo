package main

import (
	"gomongo/src/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	server := server.NewServer()

	server.Run()
}
