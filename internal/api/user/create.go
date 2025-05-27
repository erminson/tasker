package user

import (
	"net/http"

	"github.com/erminson/tasker/internal/rest"
)

type CreateUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (s *Implementation) CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := rest.DecodeJson[CreateUser](r.Body)
	if err != nil {
		rest.BadRequest(w, "invalid body")
		return
	}

	err = s.userService.Create(r.Context(), user.Login, user.Password)
	if err != nil {
		rest.Internal(w, err)
		return
	}

	rest.OK(w, nil)
}
