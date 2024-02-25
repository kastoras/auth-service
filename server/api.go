package server

import (
	"auth-service/helpers"
	"net/http"

	"github.com/gorilla/mux"
)

func NewAPIServer() *APIServer {
	httpClient := &http.Client{}
	createCacheClient()
	return &APIServer{
		addres: helpers.GetEnvParam("SERVICE_PORT", "3030"),
		client: &Client{
			httpClient: httpClient,
		},
		keycloack: &KeycloackClient{
			Host:            helpers.GetEnvParam("KEYCKLOACK_HOST", ""),
			ClientID:        helpers.GetEnvParam("KEYCKLOACK_CLIENT_ID", ""),
			ClientSecret:    helpers.GetEnvParam("KEYCKLOACK_CLIENT_SECRET", ""),
			ClientGrandType: helpers.GetEnvParam("KEYCKLOACK_GRAND_TYPE", ""),
		},
		cache: &Cache{},
	}
}

func (s *APIServer) Start(router *mux.Router) {
	http.ListenAndServe(s.addres, router)
}

func (s *APIServer) Shutdown() {
	closeCacheClient()
}

func (s *APIServer) Client() *http.Client {
	return s.client.httpClient
}

func (s *APIServer) KCClient() *KeycloackClient {
	return s.keycloack
}

func (s *APIServer) Cache() *Cache {
	return s.cache
}
