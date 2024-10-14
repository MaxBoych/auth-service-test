package userDBRepository

import (
	"auth-service-test/internal/domain"
	"context"
)

// GetOne TODO
func (r *UserRepository) GetOne(_ context.Context, _ domain.Username) (domain.User, error) {
	return domain.User{}, nil
}
