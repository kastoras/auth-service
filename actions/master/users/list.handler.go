package master_users

import (
	"auth-service/helpers"
	"auth-service/server"
	"net/http"
	"sync"
)

var (
	muc  *MasterUsersController
	once sync.Once
)

// Get users
//
// @Summary      Get users
// @Description  Get users with pagination in for current realm
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        limit query  int false "Limit"
// @Param        page query  int false "Page"
// @Success      201  {object}  interface{}
// @Failure      404  {object}  interface{}
// @Failure      409  {object}  interface{}
// @Failure      500  {object}  interface{}
// @Router       /users [get]
func Handle(server *server.APIServer) http.HandlerFunc {

	muc := getInstance()
	if muc.server == nil {
		muc.server = server
	}

	return func(w http.ResponseWriter, r *http.Request) {

		limit, page := parseQueryPagination(r)

		pagination, err := muc.CalculatePagination(limit, page)
		if err != nil {
			muc.FailedResp(&w, http.StatusInternalServerError, err.Error())
			return
		}

		users, err := muc.List(pagination)
		if err != nil {
			muc.FailedResp(&w, http.StatusInternalServerError, err.Error())
			return
		}

		uresp := UserResponse{
			Users: users,
			Pagination: PaginationResp{
				TotalEntries:   pagination.TotalEntries,
				CurrentEntries: len(users),
				CurrentPage:    pagination.CurrentPage,
				TotalPages:     pagination.TotalPages,
				Limit:          pagination.Limit,
			},
		}

		muc.SuccessWithData(&w, &uresp)
	}
}

func parseQueryPagination(r *http.Request) (int, int) {

	limit, err := helpers.ParseQueryVarToInt(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	page, err := helpers.ParseQueryVarToInt(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	return limit, page
}

func getInstance() *MasterUsersController {
	once.Do(func() {
		muc = &MasterUsersController{}
	})
	return muc
}
