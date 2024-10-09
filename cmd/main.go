package main

import (
	"log"
	"net/http"

	"github.com/Bhushan9001/GO_CRUD/config"
	"github.com/Bhushan9001/GO_CRUD/internal/routes"
	"github.com/rs/cors"
)

func main() {

	config.ConnectDB()

	router := routes.Routes()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Add your frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Wrap the router with the CORS middleware
	handler := c.Handler(router)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))

}
