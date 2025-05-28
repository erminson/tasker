package user

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/erminson/tasker/internal/rest"
)

type CompleteTask struct {
	Name string `json:"name"`
}

func (s *Implementation) CompleteTask(w http.ResponseWriter, r *http.Request) {
	completeTask, err := rest.DecodeJson[CompleteTask](r.Body)
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

	err = ValidateUser(r.Context(), id)
	if err != nil {
		rest.Forbidden(w, err.Error())
		return
	}

	task, err := s.taskService.GetTask(r.Context(), completeTask.Name)
	if err != nil {
		rest.Internal(w, err)
		return
	}

	err = s.userService.UpdatePoints(r.Context(), id, task.Points)
	if err != nil {
		rest.Internal(w, err)
		return
	}

	rest.OK(w, nil)
}
