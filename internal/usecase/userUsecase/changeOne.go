package userUsecase

import (
	"auth-service-test/internal/domain"
	"auth-service-test/pkg/errs"
	"context"
	"github.com/pkg/errors"
)

func (u *UserUsecase) ChangeOne(ctx context.Context, upd domain.UpdatedUserData) error {
	spanName := "UserUsecase.ChangeOne"

	user, err := u.userRepo.GetOne(ctx, upd.Username)
	if err != nil && !errors.Is(err, errs.UserNotFound) {
		return errors.Wrap(err, spanName)
	}

	err = u.userRepo.UpdateOne(ctx, user, upd)
	if err != nil {
		return errors.Wrap(err, spanName)
	}

	return nil
}
