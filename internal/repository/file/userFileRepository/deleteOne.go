package userFileRepository

import (
	"auth-service-test/internal/domain"
	"context"
)

// DeleteOne TODO
func (r *UserRepository) DeleteOne(_ context.Context, _ domain.Username) error {
	return nil
}
