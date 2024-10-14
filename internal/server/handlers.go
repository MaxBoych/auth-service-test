package server

import (
	"auth-service-test/internal/handlers/userHTTPHandlers"
	"auth-service-test/internal/middleware"
	"auth-service-test/internal/repository/db/userDBRepository"
	"auth-service-test/internal/repository/file/userFileRepository"
	"auth-service-test/internal/repository/inmemory/userInMemoryRepository"
	"auth-service-test/internal/usecase/userUsecase"
	"auth-service-test/pkg/consts"
	"auth-service-test/pkg/errs"
)

func (s *Server) MapHandlers() error {
	var userRepo userUsecase.UserRepository
	switch s.cfg.DB.Type {
	case consts.InMemoryType:
		userRepo = userInMemoryRepository.New()
	case consts.FileType:
		userRepo = userFileRepository.New()
	case consts.DBType:
		userRepo = userDBRepository.New()
	default:
		return errs.InvalidDBType
	}

	userUC := userUsecase.New(userRepo)
	userHandlers := userHTTPHandlers.New(userUC)
	mwManager := middleware.NewMiddlewareManager(userUC)

	userGroup := s.fiber.Group("user")
	userHTTPHandlers.MapUserRoutes(userGroup, userHandlers, mwManager)

	return nil
}
