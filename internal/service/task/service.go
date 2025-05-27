package task

import (
	"context"

	"github.com/erminson/tasker/internal/model"
	repo "github.com/erminson/tasker/internal/repository"
	"github.com/erminson/tasker/internal/service"
)

type taskService struct {
	repo repo.TaskRepository
}

func NewService(repo repo.TaskRepository) service.TaskService {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) GetTask(ctx context.Context, name string) (*model.Task, error) {
	repoTask, err := s.repo.GetTask(ctx, name)
	if err != nil {
		return nil, err
	}

	return &model.Task{
		Name:   repoTask.Name,
		Points: repoTask.Points,
	}, nil
}
