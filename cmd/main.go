// @title My Transactions API
// @description This is a sample API for managing transactions.
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
package main

import (
	"log"

	"TraiveTest/internal"
)

func main() {
	// Setup the server and configuration
	appServer, err := internal.SetupServer()
	if err != nil {
		log.Fatalf("Failed to setup server: %v", err)
	}

	// Run the server on a specified port
	port := 8080 // Change this to your desired port
	appServer.Run(port)
}
