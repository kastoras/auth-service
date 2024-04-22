package master_users

import (
	"encoding/json"
	"net/http"
)

func (muc *MasterUsersController) FailedResp(w *http.ResponseWriter, errorCode int, errorMsg string) {
	writer := *w
	writer.WriteHeader(errorCode)
	writer.Write([]byte(errorMsg))
}

func (muc *MasterUsersController) SuccessWithData(w *http.ResponseWriter, resp *UserResponse) {
	writer := *w
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(resp)
}
