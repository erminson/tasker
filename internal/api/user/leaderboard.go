package user

import (
	"net/http"

	"github.com/erminson/tasker/internal/model"
	"github.com/erminson/tasker/internal/rest"
)

type Leaderboard struct {
	Users []model.User `json:"users"`
}

func (s *Implementation) LeaderBoard(w http.ResponseWriter, r *http.Request) {
	// TODO: count add to config
	count := 10
	user, err := s.userService.GetTopUsers(r.Context(), count)
	if err != nil {
		rest.Internal(w, err)
		return
	}

	rest.OK(w, Leaderboard{
		Users: user,
	})
}
