package userUsecase

import (
	"auth-service-test/internal/domain"
	"auth-service-test/pkg/errs"
	"context"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserUsecase) RegisterOne(ctx context.Context, user domain.User) error {
	spanName := "UserUsecase.RegisterOne"

	// check user existing
	existedUser, err := u.userRepo.GetOne(ctx, user.Username)
	if err != nil && !errors.Is(err, errs.UserNotFound) {
		return errors.Wrap(err, spanName)
	}
	if !existedUser.IsEmpty() {
		return errors.Wrap(errs.UserAlreadyExists, spanName)
	}

	// generate password hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, spanName)
	}
	user.Password = string(hashedPassword)

	err = u.userRepo.CreateOne(ctx, user)
	if err != nil {
		return errors.Wrap(err, spanName)
	}

	return nil
}
