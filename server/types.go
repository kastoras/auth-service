package server

import (
	"net/http"
)

type APIServer struct {
	addres    string
	client    *Client
	keycloack *KeycloackClient
}

type Client struct {
	httpClient *http.Client
}

type KeycloackClient = struct {
	Host            string
	ClientID        string
	ClientSecret    string
	ClientGrandType string
}
