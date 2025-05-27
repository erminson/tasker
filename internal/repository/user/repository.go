package user

import (
	"context"

	"github.com/erminson/tasker/internal/repository"
	database "github.com/erminson/tasker/pkg/db"
)

type repo struct {
	db database.Driver
}

func NewUserRepository(db database.Driver) repository.UserRepository {
	return &repo{
		db: db,
	}
}

func (a *repo) Count(ctx context.Context) (int, error) {
	var count int

	err := a.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (a *repo) Save(ctx context.Context, login, passHash string) error {
	_, err := a.db.ExecContext(ctx, `
		INSERT INTO users (login, password_hash)
		VALUES ($1, $2)
	`, login, passHash)

	return err
}
