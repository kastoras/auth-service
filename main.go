package main

import (
	"auth-service/routes"
	"auth-service/server"
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	server := server.NewAPIServer()

	router = routes.RegisterAPIRoutes(router, server)

	fmt.Println("Server start...")
	server.Start(router)
}
