package user

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/erminson/tasker/internal/rest"
)

type UserInfo struct {
	Name      string `json:"name"`
	Login     string `json:"login"`
	Points    int64  `json:"points"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (s *Implementation) UserInfo(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		rest.BadRequest(w, "invalid id")
		return
	}

	err = ValidateUser(r.Context(), id)
	if err != nil {
		rest.Forbidden(w, err.Error())
		return
	}

	user, err := s.userService.GetUserById(r.Context(), id)
	if err != nil {
		rest.Internal(w, err)
		return
	}

	rest.OK(w, UserInfo{
		Name:      user.Name,
		Login:     user.Login,
		Points:    user.Points,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}
