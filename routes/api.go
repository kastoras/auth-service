package routes

import (
	"auth-service/actions"
	"auth-service/actions/login"
	master_groups "auth-service/actions/master/groups"
	master_users "auth-service/actions/master/users"
	"auth-service/actions/register"
	"auth-service/actions/token"
	"auth-service/server"

	"github.com/gorilla/mux"

	_ "auth-service/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type HandlerContext struct {
	*server.APIServer
}

func RegisterAPIRoutes(router *mux.Router, s *server.APIServer) *mux.Router {

	router.Use(s.ClientSetUp)

	router.HandleFunc("/health", actions.HandleAPIHealth(s)).Methods("GET")

	//public
	router.HandleFunc("/login", login.HandleLogin(s)).Methods("POST")
	router.HandleFunc("/groups", master_groups.Handle(s)).Methods("GET")
	router.HandleFunc("/register", register.Handle(s)).Methods("POST")

	//client
	router.HandleFunc("/token", token.Handle(s)).Methods("POST")

	//admin
	router.HandleFunc("/users", master_users.Handle(s)).Methods("GET")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return router
}
