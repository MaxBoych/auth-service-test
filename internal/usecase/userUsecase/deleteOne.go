package userUsecase

import (
	"auth-service-test/internal/domain"
	"context"
	"github.com/pkg/errors"
)

func (u *UserUsecase) DeleteOne(ctx context.Context, username domain.Username) error {
	spanName := "UserUsecase.DeleteOne"

	err := u.userRepo.DeleteOne(ctx, username)
	if err != nil {
		return errors.Wrap(err, spanName)
	}

	return nil
}
