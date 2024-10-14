package userInMemoryRepository

import (
	"auth-service-test/internal/domain"
	"context"
)

func (r *UserRepository) CreateOne(_ context.Context, user domain.User) error {
	r.users.Store(user.Username, user)

	return nil
}
