package token

import "auth-service/server"

type TokenController struct {
	Server *server.APIServer
}

type JWTContent struct{}

type KeycloakTokenIntrospectPayload struct {
	ClientID     string `json:"client_id"`
	Token        string `json:"token"`
	ClientSecret string `json:"client_secret"`
}

type KeycloakResp struct {
	Active bool  `json:"active"`
	Expire int64 `json:"exp"`
}
