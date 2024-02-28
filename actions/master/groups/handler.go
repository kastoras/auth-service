package master_groups

import (
	"auth-service/server"
	"encoding/json"
	"net/http"
	"sync"
)

var (
	mgc  *MasterGroupsController
	once sync.Once
)

// Get User Groups
//
// @Summary      Return application user groups
// @Description  Returns the groups a user can be assigned
// @Tags         managment
// @Accept       json
// @Produce      json
// @Success      200  {object}  master_groups.Response
// @Failure      400  {object}  interface{}
// @Failure      404  {object}  interface{}
// @Failure      500  {object}  interface{}
// @Router       /groups [get]
func Handle(server *server.APIServer) http.HandlerFunc {

	mgc := getInstance()
	if mgc.server == nil {
		mgc.server = server
	}

	return func(w http.ResponseWriter, r *http.Request) {

		groups, err := mgc.keycloackGroups()
		if err != nil {
			failedResp(&w, 400, err.Error())
			return
		}

		res := &Response{
			Groups: groups,
		}

		successResp(&w, 200, res)
	}
}

func getInstance() *MasterGroupsController {
	once.Do(func() {
		mgc = &MasterGroupsController{}
	})
	return mgc
}

func failedResp(w *http.ResponseWriter, errorCode int, errorMsg string) {
	writer := *w
	writer.WriteHeader(errorCode)
	writer.Write([]byte(errorMsg))
}

func successResp(w *http.ResponseWriter, code int, resp *Response) {
	writer := *w
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(resp)
}
