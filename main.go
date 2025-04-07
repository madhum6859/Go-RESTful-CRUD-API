package main

import (
	"github.com/trae/go-restful-crud-api/routes"
)

func main() {
	// Initialize the router
	router := routes.SetupRouter()

	// Start the server
	router.Run(":8080")
}