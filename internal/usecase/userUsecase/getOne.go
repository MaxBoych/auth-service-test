package userUsecase

import (
	"auth-service-test/internal/domain"
	"context"
	"github.com/pkg/errors"
)

func (u *UserUsecase) GetOne(ctx context.Context, username domain.Username) (domain.User, error) {
	spanName := "UserUsecase.GetOne"

	user, err := u.userRepo.GetOne(ctx, username)
	if err != nil {
		return domain.User{}, errors.Wrap(err, spanName)
	}

	return user, nil
}
