package user

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/erminson/tasker/internal/rest"
)

type UpdateUser struct {
	Name string `json:"name"`
}

func (s *Implementation) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user, err := rest.DecodeJson[UpdateUser](r.Body)
	if err != nil {
		rest.BadRequest(w, "invalid body")
		return
	}

	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		rest.BadRequest(w, "invalid id")
		return
	}

	err = s.userService.UpdateName(r.Context(), id, user.Name)
	if err != nil {
		rest.Internal(w, err)
		return
	}

	rest.OK(w, nil)
}
