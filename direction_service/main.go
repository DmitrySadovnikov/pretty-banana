package main

import (
	"direction_service/app/controllers"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	godotenv.Load()
}

func main() {
	log.Printf("Hello")
	server := controllers.NewServer()

	go func() {
		log.Printf("Server running on port " + os.Getenv("PORT") + "...")
		err := server.Server.ListenAndServe()
		if err != nil {
			log.Printf("Error: %v", err)
		}
	}()

	server.WaitShutdown()
	log.Printf("Server shuted down")
}
