package login

import (
	"auth-service/server"
	"encoding/json"
	"net/http"
)

func HandleLogin(server *server.APIServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload LoginPayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			failedResp(&w, 404, "Invalid payload")
			return
		}

		kresp, err := keycloackLogin(server, payload)
		if err != nil {
			failedResp(&w, 400, err.Error())
			return
		}

		res := &LoginResp{
			AccessToken: kresp.AccessToken,
		}

		successResp(&w, 200, res)
	}
}

func failedResp(w *http.ResponseWriter, errorCode int, errorMsg string) {
	writer := *w
	writer.WriteHeader(errorCode)
	writer.Write([]byte(errorMsg))
}

func successResp(w *http.ResponseWriter, code int, resp *LoginResp) {
	writer := *w
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(resp)
}
