package actions

import (
	"auth-service/helpers"
	"auth-service/server"
	"encoding/json"
	"net/http"
)

type HealthRes struct {
	ServerStatus string `json:"server-status"`
	CacheStatus  string `json:"cache-status"`
	Version      string `json:"version"`
}

func HandleAPIHealth(server *server.APIServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		currentAPIVersion := helpers.GetEnvParam("API_VERSION", "")

		cacheStatus := server.Cache().Ping()

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)

		res := &HealthRes{
			ServerStatus: "running",
			CacheStatus:  cacheStatus,
			Version:      currentAPIVersion,
		}
		json.NewEncoder(w).Encode(res)
	}
}
