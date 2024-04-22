package register

import "auth-service/server"

type RegisterUserController struct {
	server *server.APIServer
}

type RegisterPayload struct {
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Password  string   `json:"password"`
	Groups    []string `json:"groups"`
}

type KRegisterPayload struct {
	Username        string   `json:"username"`
	Email           string   `json:"email"`
	Firstname       string   `json:"firstName"`
	Lastname        string   `json:"lastName"`
	RequiredActions []string `json:"requiredActions"`
	EmailVerified   bool     `json:"emailVerified"`
	Groups          []string `json:"groups"`
	Enabled         bool     `json:"enabled"`
}

type KPasswordPayload struct {
	Temporary bool   `json:"temporary"`
	Type      string `json:"type"`
	Value     string `json:"value"`
}
