package task

import (
	"context"

	"github.com/erminson/tasker/internal/repository"
	"github.com/erminson/tasker/internal/repository/model"
	db "github.com/erminson/tasker/pkg/db"
)

type repo struct {
	db db.Driver
}

func NewRepository(db db.Driver) repository.TaskRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) GetTask(ctx context.Context, name string) (*model.Task, error) {
	q := `
		SELECT name, points FROM tasks WHERE name = $1
	`

	var task model.Task
	err := r.db.QueryRowContext(ctx, q, name).Scan(&task.Name, &task.Points)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
