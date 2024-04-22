package register

import (
	"auth-service/server"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var (
	ruc  *RegisterUserController
	once sync.Once
)

// Register new user
//
// @Summary      Creates a new user
// @Description  Creates a new user in user group client
// @Tags         authentication
// @Accept       json
// @Produce      json
// @Param        request body  register.RegisterPayload true "User Data"
// @Success      201  {object}  interface{}
// @Failure      404  {object}  interface{}
// @Failure      409  {object}  interface{}
// @Failure      500  {object}  interface{}
// @Router       /register [post]
func Handle(server *server.APIServer) http.HandlerFunc {

	ruc = getInstance()
	if ruc.server == nil {
		ruc.server = server
	}

	return func(w http.ResponseWriter, r *http.Request) {

		var payload RegisterPayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			failedResp(&w, http.StatusBadRequest, "Invalid payload")
			return
		}

		statusCode, err := ruc.Register(payload)
		if err != nil {
			failedResp(&w, statusCode, err.Error())
			return
		}

		successResp(&w)
	}
}

func getInstance() *RegisterUserController {
	once.Do(func() {
		ruc = &RegisterUserController{}
	})
	return ruc
}

func failedResp(w *http.ResponseWriter, errorCode int, errorMsg string) {

	fmt.Printf("Error on user regiser. Service status code: %v, Service error msg: %v \n", errorCode, errorMsg)
	var respcode int

	switch errorCode {
	case http.StatusConflict:
		errorMsg = "User name already exists!"
		respcode = http.StatusConflict
	case http.StatusBadRequest:
		errorMsg = "Invalid data"
		respcode = http.StatusBadRequest
	default:
		errorMsg = "Internal server error"
		respcode = http.StatusInternalServerError
	}

	writer := *w
	writer.WriteHeader(respcode)
	writer.Write([]byte(errorMsg))
}

func successResp(w *http.ResponseWriter) {
	writer := *w
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
}
