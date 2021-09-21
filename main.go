package main

import "gomongo/src/server"

func main() {
	server := server.NewServer()

	server.Run()
}
