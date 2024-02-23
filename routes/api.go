package routes

import (
	"auth-service/actions"
	"auth-service/actions/login"
	"auth-service/actions/token"
	"auth-service/server"

	"github.com/gorilla/mux"

	_ "auth-service/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type HandlerContext struct {
	*server.APIServer
}

func RegisterAPIRoutes(router *mux.Router, server *server.APIServer) *mux.Router {

	router.HandleFunc("/health", actions.HandleAPIHealth(server)).Methods("GET")

	router.HandleFunc("/login", login.HandleLogin(server)).Methods("POST")
	router.HandleFunc("/token", token.Handle(server)).Methods("POST")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return router
}
