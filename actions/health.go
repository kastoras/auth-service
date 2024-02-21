package actions

import (
	"auth-service/helpers"
	"auth-service/server"
	"encoding/json"
	"net/http"
)

type HealthRes struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func HandleAPIHealth(server *server.APIServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)

		currentAPIVersion := helpers.GetEnvParam("API_VERSION", "")

		res := &HealthRes{
			Status:  "running",
			Version: currentAPIVersion,
		}
		json.NewEncoder(w).Encode(res)
	}
}
