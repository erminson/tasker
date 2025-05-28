package user

import (
	"context"
	"fmt"

	"github.com/erminson/tasker/internal/model"
	repo "github.com/erminson/tasker/internal/repository"
	"github.com/erminson/tasker/internal/service"
	"github.com/erminson/tasker/pkg/crypto"
)

const ADMIN = "admin"

type userService struct {
	repo repo.UserRepository
}

func NewService(repo repo.UserRepository) service.UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Create(ctx context.Context, login, password string) error {
	return s.repo.Save(ctx, login, crypto.BCrypto(password))
}

func (s *userService) UpdateName(ctx context.Context, id int64, name string) error {
	if name == "" {
		return fmt.Errorf("name is empty")
	}

	return s.repo.UpdateName(ctx, id, name)
}

func (s *userService) UpdatePoints(ctx context.Context, id int64, points int64) error {
	if points < 0 {
		return fmt.Errorf("points cannot be negative")
	}

	return s.repo.UpdatePoints(ctx, id, points)
}

func (s *userService) GetTopUsers(ctx context.Context, count int) ([]model.User, error) {
	repoUsers, err := s.repo.GetTopUsers(ctx, count)
	if err != nil {
		return nil, err
	}

	users := make([]model.User, 0, len(repoUsers))
	for _, u := range repoUsers {
		// TODO: use convertor
		if u.Login == ADMIN {
			continue
		}

		name := u.Login
		if u.Name != nil && *u.Name != "" {
			name = *u.Name
		}

		user := model.User{
			Points: u.Points,
			Name:   name,
		}

		users = append(users, user)
	}

	return users, nil
}

func (s *userService) Referrer(ctx context.Context, userID, referrerID int64) error {
	return s.repo.Referrer(ctx, userID, referrerID)
}

func (s *userService) ValidateUser(ctx context.Context, login, password string) (int64, error) {
	repoUser, err := s.repo.GetUserByLogin(ctx, login)
	if err != nil {
		return -1, err
	}

	return repoUser.Id, crypto.CheckPasswordHash(repoUser.PassHash, password)
}

func (s *userService) GetUserById(ctx context.Context, id int64) (*model.UserInfo, error) {
	repoUser, err := s.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	name := repoUser.Login
	if repoUser.Name != nil && *repoUser.Name != "" {
		name = *repoUser.Name
	}

	user := model.UserInfo{
		Name:      name,
		Login:     repoUser.Login,
		Points:    repoUser.Points,
		CreatedAt: repoUser.CreatedAt,
		UpdatedAt: repoUser.UpdatedAt,
	}

	return &user, nil
}
