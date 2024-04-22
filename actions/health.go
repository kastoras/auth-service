package actions

import (
	"auth-service/helpers"
	master_token "auth-service/helpers/masterrealm/token"
	"auth-service/server"
	"encoding/json"
	"net/http"
)

type HealthRes struct {
	ServerStatus string `json:"server-status"`
	CacheStatus  string `json:"cache-status"`
	AdminToken   string `json:"admin-token"`
	Version      string `json:"version"`
}

// API's Health Check
//
// @Summary      Checks the health of the API
// @Description  Check's the health of the api and all its components
// @Tags         information
// @Produce      json
// @Success      200  {object}  actions.HealthRes
// @Failure      500  {object}  interface{}
// @Router       /health [get]
func HandleAPIHealth(server *server.APIServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		currentAPIVersion := helpers.GetEnvParam("API_VERSION", "")

		cacheStatus := server.Cache().Ping()

		admintoken, err := master_token.Get(server)
		if err != nil {
			admintoken = "no admin connection..."
		}
		if admintoken != "" {
			admintoken = "admin connected!"
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)

		res := &HealthRes{
			ServerStatus: "running",
			CacheStatus:  cacheStatus,
			AdminToken:   admintoken,
			Version:      currentAPIVersion,
		}
		json.NewEncoder(w).Encode(res)
	}
}
