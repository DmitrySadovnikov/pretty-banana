package main

import (
	"log"
	"direction_service/app"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	log.Printf("Hello")
	go app.StartHTTPServer()
	app.StartGRPCServer()
}
