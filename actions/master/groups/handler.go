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
// @Produce      json
// @Success      200  {object}  master_groups.Response
// @Success      204  {object}  interface{}
// @Failure      500  {object}  interface{}
// @Router       /groups [get]
func Handle(server *server.APIServer) http.HandlerFunc {

	mgc := getInstance()
	if mgc.server == nil {
		mgc.server = server
	}

	return func(w http.ResponseWriter, r *http.Request) {

		groups, err := mgc.getFromCache()

		if err == nil {
			successResp(&w, &groups)
			return
		}

		groups, err = mgc.keycloackGroups()
		if err != nil {
			failedResp(&w, http.StatusInternalServerError, err.Error())
			return
		}

		successResp(&w, &groups)
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

func successResp(w *http.ResponseWriter, groups *[]RealmGroup) {

	if groups == nil || len(*groups) == 0 {
		successNoContent(w)
		return
	}

	resp := &Response{
		Groups: *groups,
	}

	successWithData(w, resp)
}

func successNoContent(w *http.ResponseWriter) {
	writer := *w
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNoContent)
}

func successWithData(w *http.ResponseWriter, resp *Response) {
	writer := *w
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(resp)
}
