package repository

import "context"

type UserRepository interface {
	Count(ctx context.Context) (int, error)
	Save(ctx context.Context, login, passHash string) error
}
