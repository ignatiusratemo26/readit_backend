package main

import (
	"log"
	"net/http"

	"readit_backend/routes"

	"readit_backend/data"

	"github.com/gorilla/handlers"
)

func main() {
	data.InitMongo()

	r := routes.RegisterRoutes()

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "Accept", "X-Requested-With", "Origin"}),
		handlers.AllowCredentials(),
	)

	log.Println("Server running on:4000")
	http.ListenAndServe(":4000", cors(r))

}
