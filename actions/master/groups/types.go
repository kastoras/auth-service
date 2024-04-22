package master_groups

import "auth-service/server"

type MasterGroupsController struct {
	server *server.APIServer
}

type KResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type RealmGroup struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Groups []RealmGroup `json:"groups"`
}
