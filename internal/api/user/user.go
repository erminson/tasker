package user

import (
	"context"
	"fmt"

	"github.com/erminson/tasker/internal/rest"
	"github.com/erminson/tasker/internal/service"
)

const ID = "user_id"

type Implementation struct {
	userService service.UserService
	taskService service.TaskService
	authMid     *rest.AuthMiddleware
}

func NewApi(userService service.UserService, taskService service.TaskService, authMid *rest.AuthMiddleware) *Implementation {
	return &Implementation{
		userService: userService,
		taskService: taskService,
		authMid:     authMid,
	}
}

func ValidateUser(ctx context.Context, id int64) error {
	userId, ok := ctx.Value(ID).(int64)
	if !ok {
		return fmt.Errorf("invalid user id")
	}

	if id != userId {
		return fmt.Errorf("you can't complete task for another user")
	}

	return nil
}
