package userUsecase

import (
	"auth-service-test/internal/domain"
	"auth-service-test/pkg/errs"
	"auth-service-test/pkg/logger"
	"context"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserUsecase) Authenticate(
	ctx context.Context,
	username domain.Username,
	password string,
) error {
	spanName := "UserUsecase.Authenticate"

	user, err := u.GetOne(ctx, username)
	if err != nil {
		return errors.Wrap(err, spanName)
	}

	if !isPasswordValid(user.Password, password) {
		return errors.Wrap(errs.IncorrectPassword, spanName)
	}

	return nil
}

func isPasswordValid(hashedPasswd, passwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPasswd), []byte(passwd)); err != nil {
		logger.Error(err.Error())
		return false
	}
	return true
}
