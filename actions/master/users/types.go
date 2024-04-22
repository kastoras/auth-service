package master_users

import (
	"auth-service/server"
)

type MasterUsersController struct {
	server *server.APIServer
}

type KUser struct {
	ID            string `json:"id"`
	Created       int64  `json:"createdTimestamp"`
	Username      string `json:"username"`
	Enabled       bool   `json:"enabled"`
	EmailVerified bool   `json:"emailVerified"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	Attributes    map[string]interface{}
}

type User struct {
	ID            string                 `json:"id"`
	Created       int64                  `json:"creted"`
	Username      string                 `json:"username"`
	Enabled       bool                   `json:"enabled"`
	EmailVerified bool                   `json:"verified"`
	FirstName     string                 `json:"firstname"`
	LastName      string                 `json:"lastname"`
	Email         string                 `json:"email"`
	Attributes    map[string]interface{} `json:"attrs"`
}

type UserResponse struct {
	Users      []User         `json:"users"`
	Pagination PaginationResp `json:"pagination"`
}

type QueryParamsPagination struct {
	Limit int
	Page  int
}

type Pagination struct {
	TotalEntries int `json:"total_entries"`
	CurrentPage  int `json:"current_page"`
	TotalPages   int `json:"total_pages"`
	Limit        int `json:"limit"`
	Offset       int `json:"offset"`
}

type PaginationResp struct {
	TotalEntries   int `json:"total_entries"`
	CurrentEntries int `json:"page_entries"`
	CurrentPage    int `json:"current_page"`
	TotalPages     int `json:"total_pages"`
	Limit          int `json:"limit"`
}
