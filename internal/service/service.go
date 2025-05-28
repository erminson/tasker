package service

import (
	"context"

	"github.com/erminson/tasker/internal/model"
)

type UserService interface {
	Create(ctx context.Context, login, password string) error
	UpdateName(ctx context.Context, id int64, name string) error
	UpdatePoints(ctx context.Context, id int64, points int64) error
	GetTopUsers(ctx context.Context, count int) ([]model.User, error)
	Referrer(ctx context.Context, userID, referrerID int64) error
	ValidateUser(ctx context.Context, login, password string) (int64, error)
	GetUserById(ctx context.Context, id int64) (*model.UserInfo, error)
}

type TaskService interface {
	GetTask(ctx context.Context, name string) (*model.Task, error)
}
