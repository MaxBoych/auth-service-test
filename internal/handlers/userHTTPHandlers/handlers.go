package userHTTPHandlers

import (
	"auth-service-test/internal/domain"
	"auth-service-test/internal/usecase/userUsecase"
	"context"
	"github.com/gofiber/fiber/v2"
)

type (
	Handlers interface {
		RegisterOne() fiber.Handler
		GetOne() fiber.Handler
		GetOneOwn() fiber.Handler
		ChangeOne() fiber.Handler
		DeleteOne() fiber.Handler
	}

	UserUsecase interface {
		RegisterOne(ctx context.Context, user domain.User) error
		GetOne(ctx context.Context, username domain.Username) (domain.User, error)
		ChangeOne(ctx context.Context, upd domain.UpdatedUserData) error
		DeleteOne(ctx context.Context, username domain.Username) error
	}
)

var _ Handlers = (*UserHandlers)(nil)
var _ UserUsecase = (*userUsecase.UserUsecase)(nil)

type UserHandlers struct {
	userUC UserUsecase
}

func New(
	userUC UserUsecase,
) *UserHandlers {
	return &UserHandlers{
		userUC: userUC,
	}
}
