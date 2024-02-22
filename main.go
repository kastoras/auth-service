package main

import (
	"auth-service/routes"
	"auth-service/server"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	server := server.NewAPIServer()

	router = routes.RegisterAPIRoutes(router, server)

	server.Start(router)
}
