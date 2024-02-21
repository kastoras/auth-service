package main

import (
	"auth-service/routes"
	"auth-service/server"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	address := ":3030"

	server := server.NewAPIServer(address)
	router = routes.RegisterAPIRoutes(router, server)

	server.Run(router)
}
