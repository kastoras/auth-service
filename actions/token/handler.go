package token

import (
	"auth-service/helpers"
	"auth-service/server"
	"net/http"
	"sync"
)

var (
	tc   *TokenController
	once sync.Once
)

// Token Check
//
// @Summary      Checks if token is valid
// @Description  Check if Authentication Bearer token is active
// @Tags         authentication
// @Produce      json
// @Success      200  {object}   interface{}
// @Failure      401  {object}  interface{}
// @Failure      404  {object}  interface{}
// @Failure      500  {object}  interface{}
// @Router       /token [post]
func Handle(server *server.APIServer) http.HandlerFunc {

	tc := getInstance()
	if tc.Server == nil {
		tc.Server = server
	}

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

		err = tc.checkTokenValidy(&jwtInfo, jwt)
		if err != nil {
			failedResp(&w, http.StatusUnauthorized, "")
			return
		}

		successResp(&w)
	}
}

func getInstance() *TokenController {
	once.Do(func() {
		tc = &TokenController{}
	})
	return tc
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
