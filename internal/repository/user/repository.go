package user

import (
	"context"

	"github.com/erminson/tasker/internal/repository"
	"github.com/erminson/tasker/internal/repository/model"
	db "github.com/erminson/tasker/pkg/db"
)

type repo struct {
	db db.Driver
}

func NewRepository(db db.Driver) repository.UserRepository {
	return &repo{
		db: db,
	}
}

func (a *repo) Count(ctx context.Context) (int, error) {
	q := `SELECT COUNT(*) FROM users`

	var count int
	err := a.db.QueryRowContext(ctx, q).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (a *repo) Save(ctx context.Context, login, passHash string) error {
	q := `
		INSERT INTO users (login, password_hash)
		VALUES ($1, $2)
	`

	_, err := a.db.ExecContext(ctx, q, login, passHash)

	return err
}

func (a *repo) UpdateName(ctx context.Context, id int64, name string) error {
	q := `
    	UPDATE users
    	SET name = $1
    	WHERE user_id = $2
	`

	_, err := a.db.ExecContext(ctx, q, name, id)
	if err != nil {
		return err
	}

	return nil
}

func (a *repo) UpdatePoints(ctx context.Context, id int64, points int64) error {
	q := `
    	UPDATE users
    	SET points = points + $1
    	WHERE user_id = $2
	`

	_, err := a.db.ExecContext(ctx, q, points, id)
	if err != nil {
		return err
	}

	return nil
}

func (a *repo) GetTopUsers(ctx context.Context, count int) ([]model.User, error) {
	q := `
		SELECT user_id, login, name, points, created_at, updated_at
		FROM users
		ORDER BY points DESC
		LIMIT $1
	`

	rows, err := a.db.QueryContext(ctx, q, count)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var users []model.User
	for rows.Next() {
		var u model.User
		err = rows.Scan(&u.Id, &u.Login, &u.Name, &u.Points, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}
