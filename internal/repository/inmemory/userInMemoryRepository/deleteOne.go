package userInMemoryRepository

import (
	"auth-service-test/internal/domain"
	"context"
)

func (r *UserRepository) DeleteOne(_ context.Context, username domain.Username) error {
	r.users.Delete(username)

	return nil
}
