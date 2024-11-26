package main

import (
	"log"
	"obsidian/practice/router"
	"os"
)

func main() {
	// Get the port from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8000 for local development
	}

	r, err := router.NewRouter()
	if err != nil {
		log.Fatal(err)
	}
	r.Run(":" + port) // Use the dynamic port
}
