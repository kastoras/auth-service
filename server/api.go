package server

import (
	"auth-service/helpers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewAPIServer(addres string) *APIServer {
	httpClient := &http.Client{}
	return &APIServer{
		addres: addres,
		client: &Client{
			httpClient: httpClient,
		},
		keycloack: &KeycloackClient{
			Host:            helpers.GetEnvParam("KEYCKLOACK_HOST", ""),
			ClientID:        helpers.GetEnvParam("KEYCKLOACK_CLIENT_ID", ""),
			ClientSecret:    helpers.GetEnvParam("KEYCKLOACK_CLIENT_SECRET", ""),
			ClientGrandType: helpers.GetEnvParam("KEYCKLOACK_GRAND_TYPE", ""),
		},
	}
}

func (s *APIServer) Run(router *mux.Router) {
	http.ListenAndServe(s.addres, router)
	fmt.Println("Server started...")
}

func (s *APIServer) Client() *http.Client {
	return s.client.httpClient
}

func (s *APIServer) KCClient() *KeycloackClient {
	return s.keycloack
}
