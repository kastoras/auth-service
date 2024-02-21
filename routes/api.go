package routes

import (
	"auth-service/actions"
	"auth-service/actions/login"
	"auth-service/server"

	"github.com/gorilla/mux"
)

type HandlerContext struct {
	*server.APIServer
}

func RegisterAPIRoutes(router *mux.Router, server *server.APIServer) *mux.Router {

	router.HandleFunc("/health", actions.HandleAPIHealth(server)).Methods("GET")
	router.HandleFunc("/login", login.HandleLogin(server)).Methods("POST")

	return router
}
