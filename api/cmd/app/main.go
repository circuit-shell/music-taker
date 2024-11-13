package main

import (
	"log"

	"github.com/circuit-shell/playlist-builder-back/internal/api/router"
)

func main() {
	// Get router and handle potential error
	r, err := router.SetupRouter()
	if err != nil {
		log.Fatal("Failed to setup router:", err)
	}

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
