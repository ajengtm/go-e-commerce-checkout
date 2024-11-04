package main

import (
	"fmt"
	"log"
	"main.go/internal/storage"
	"main.go/pkg/rest"
	"net/http"
)

func main() {
	// Initialize the database
	err := storage.InitDB("postgres://agussriindrawansigit:123456@localhost:5432/postgres")
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Defer closing the database connection
	defer storage.DB.Close()

	r := rest.SetupRoutes()

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", r)
}
