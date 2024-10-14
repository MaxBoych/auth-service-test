package userInMemoryRepository

import (
	"auth-service-test/internal/domain"
	"auth-service-test/pkg/errs"
	"context"
	"github.com/pkg/errors"
)

func (r *UserRepository) GetOne(_ context.Context, username domain.Username) (domain.User, error) {
	spanName := "UserRepository.GetOne"

	data, ok := r.users.Load(username)
	if !ok {
		return domain.User{}, errors.Wrap(errs.UserNotFound, spanName)
	}

	user, ok := data.(domain.User)
	if !ok {
		return domain.User{}, errors.Wrap(errs.UnknownType, spanName)
	}

	return user, nil
}
