package token

import (
	"auth-service/helpers"
	"auth-service/server"
	"net/http"
)

// Token Check
//
// @Summary      Checks if token is valid
// @Description  Check if Authentication Bearer token is active
// @Tags         authentication
// @Produce      json
// @Success      200  {array}   interface{}
// @Failure      401  {object}  interface{}
// @Failure      404  {object}  interface{}
// @Failure      500  {object}  interface{}
// @Router       /token [post]
func Handle(server *server.APIServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		jwt, err := helpers.GetJWTFromHeader(authHeader)
		if err != nil {
			failedResp(&w, http.StatusBadRequest, err.Error())
			return
		}

		jwtInfo, err := helpers.Parse(jwt)
		if err != nil {
			failedResp(&w, http.StatusBadRequest, err.Error())
			return
		}

		err = checkTokenValidy(server, &jwtInfo, jwt)
		if err != nil {
			failedResp(&w, http.StatusUnauthorized, "")
			return
		}

		successResp(&w)
	}
}

func failedResp(w *http.ResponseWriter, code int, errorMsg string) {
	writer := *w

	if code == http.StatusUnauthorized {
		writer.WriteHeader(http.StatusUnauthorized)
	} else {
		writer.WriteHeader(code)
		writer.Write([]byte(errorMsg))
	}
}

func successResp(w *http.ResponseWriter) {
	writer := *w
	writer.WriteHeader(http.StatusOK)
}
