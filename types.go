package main

import "net/http"

type KeycloakService interface {
	login(payload *KloginPayload) error
}

type KloginPayload struct {
	clientID     string
	username     string
	password     string
	grandType    string
	clientSecret string
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Client struct {
	httpClient *http.Client
}

type KLoginResp struct {
	AccessToken string `json:"access_token"`
}

type LoginResp struct {
	AccessToken string `json:"access_token"`
}
