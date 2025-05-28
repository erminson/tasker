package user

import (
	"net/http"

	"github.com/erminson/tasker/internal/rest"
)

type loginUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type userToken struct {
	Token string `json:"token"`
}

func (s *Implementation) Login(w http.ResponseWriter, r *http.Request) {
	user, err := rest.DecodeJson[loginUser](r.Body)
	if err != nil {
		rest.BadRequest(w, "invalid body")
		return
	}

	id, err := s.userService.ValidateUser(r.Context(), user.Login, user.Password)
	if err != nil {
		rest.BadRequest(w, "Invalid credentials")
		return
	}

	token, err := s.authMid.GenerateToken(id)
	if err != nil {
		rest.Internal(w, err)
		return
	}

	rest.OK(w, userToken{
		Token: token,
	})
}
