package middleware

import (
	"auth-service-test/internal/domain"
	"context"
)

type UserUsecase interface {
	Authenticate(ctx context.Context, username domain.Username, password string) error
	GetOne(ctx context.Context, username domain.Username) (domain.User, error)
}

type MiddlewareManager struct {
	userUC UserUsecase
}

func NewMiddlewareManager(userUC UserUsecase) *MiddlewareManager {
	return &MiddlewareManager{
		userUC: userUC,
	}
}
