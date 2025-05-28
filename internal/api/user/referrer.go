package user

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/erminson/tasker/internal/rest"
)

func (s *Implementation) Referrer(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		rest.BadRequest(w, "invalid id")
		return
	}

	referrerID, ok := r.Context().Value(ID).(int64) // кто приглашает
	if !ok {
		rest.BadRequest(w, "invalid user id")
	}

	referredID := id // кого пригласили

	err = s.userService.Referrer(r.Context(), referredID, referrerID)
	if err != nil {
		rest.Internal(w, err)
		return
	}

	rest.OK(w, nil)
}
