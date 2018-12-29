package app

import (
	"log"
	"os"
	c "direction_service/app/controllers"
)

func StartHTTPServer() {
	server := c.NewServer()

	go func() {
		log.Printf("HTTP server running on port " + os.Getenv("PORT") + "...")
		err := server.Server.ListenAndServe()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}()

	server.WaitShutdown()
	log.Printf("HTTP server shuted down")
}
