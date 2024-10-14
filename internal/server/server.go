package server

import (
	"auth-service-test/internal/config"
	"auth-service-test/pkg/logger"
	"github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

type Server struct {
	cfg   *config.Config
	fiber *fiber.App
}

func NewServer(
	cfg *config.Config,
) *Server {
	return &Server{
		cfg:   cfg,
		fiber: fiber.New(fiber.Config{}),
	}
}

func (s *Server) Run() error {
	s.fiber.Use(loggerMiddleware())
	s.fiber.Use(recoverMiddleware())

	if err := s.MapHandlers(); err != nil {
		return err
	}

	if !s.cfg.IsProduction {
		logger.Info("Running in develop mode")
		s.fiber.Get("/swagger/*", swagger.HandlerDefault)
	}

	go func() {
		s.fiber.Get("/health_check", func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})
		logger.Infof("Server is started on %s", s.cfg.Server.Host)

		err := s.fiber.Listen(s.cfg.Server.Host)
		if err != nil {
			logger.Errorf("Couldn't start server on %s, err=%v", s.cfg.Server.Host, err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down HTTP server...")
	if err := s.fiber.Shutdown(); err != nil {
		logger.Errorf("Couldn't shut down server, err=%v", err)
	} else {
		logger.Infof("HTTP server closed properly")
	}

	return nil
}

func loggerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger.Info(
			"Request received",
			zap.String("method", c.Method()),
			zap.String("route", c.Path()),
			zap.String("ip", c.IP()),
		)
		return c.Next()
	}
}

func recoverMiddleware() fiber.Handler {
	return recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			fullStackTrace := string(debug.Stack())
			logger.Errorf("Recovered from panic: %s", fullStackTrace)
		},
	})
}
