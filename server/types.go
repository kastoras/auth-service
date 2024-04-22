package server

import (
	"net/http"
)

type APIServer struct {
	addres    string
	client    *Client
	keycloack *KeycloackClient
	cache     *Cache
}

type Client struct {
	httpClient *http.Client
}

type KeycloackClient = struct {
	Host            string
	ClientID        string
	ClientSecret    string
	ClientGrandType string
	Realm           string
}

type Cache struct{}

type RedisCacheClient = struct {
	Host     string
	Password string
	DB       int
}

type Domain struct {
	Domain        string
	Client_id     string
	Client_secret string
	Grand_type    string
	Host          string
	Realm         string
	Active        bool
}
