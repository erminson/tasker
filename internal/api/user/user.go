package user

import (
	"github.com/erminson/tasker/internal/service"
)

type Implementation struct {
	userService service.UserService
	taskService service.TaskService
}

func NewApi(userService service.UserService, taskService service.TaskService) *Implementation {
	return &Implementation{
		userService: userService,
		taskService: taskService,
	}
}
