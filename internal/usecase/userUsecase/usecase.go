package userUsecase

import (
	"auth-service-test/internal/domain"
	"auth-service-test/internal/repository/db/userDBRepository"
	"auth-service-test/internal/repository/file/userFileRepository"
	"auth-service-test/internal/repository/inmemory/userInMemoryRepository"
	"context"
)

type UserRepository interface {
	GetOne(ctx context.Context, username domain.Username) (domain.User, error)
	CreateOne(ctx context.Context, user domain.User) error
	UpdateOne(ctx context.Context, curr domain.User, data domain.UpdatedUserData) error
	DeleteOne(ctx context.Context, username domain.Username) error
}

var _ UserRepository = (*userInMemoryRepository.UserRepository)(nil)
var _ UserRepository = (*userFileRepository.UserRepository)(nil)
var _ UserRepository = (*userDBRepository.UserRepository)(nil)

type UserUsecase struct {
	userRepo UserRepository
}

func New(
	userRepo UserRepository,
) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}
