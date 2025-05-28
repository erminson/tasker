package repository

import (
	"context"

	"github.com/erminson/tasker/internal/repository/model"
)

type UserRepository interface {
	Count(ctx context.Context) (int, error)
	Save(ctx context.Context, login, passHash string) error
	UpdateName(ctx context.Context, id int64, name string) error
	UpdatePoints(ctx context.Context, id int64, points int64) error
	GetTopUsers(ctx context.Context, count int) ([]model.User, error)
	Referrer(ctx context.Context, userID, referrerID int64) error
	GetUserByLogin(ctx context.Context, login string) (*model.LoginUser, error)
	GetUserById(ctx context.Context, id int64) (*model.User, error)
}

type TaskRepository interface {
	GetTask(ctx context.Context, name string) (*model.Task, error)
}
