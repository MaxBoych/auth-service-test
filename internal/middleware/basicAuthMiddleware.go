package middleware

import (
	"auth-service-test/internal/domain"
	"auth-service-test/pkg/consts"
	"auth-service-test/pkg/errs"
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func (mw *MiddlewareManager) BasicAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		username, password, err := parseBasicAuthHeader(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}

		if err := mw.userUC.Authenticate(c.Context(), domain.Username(username), password); err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}

		c.Locals(consts.UsernameFieldName, username)

		return c.Next()
	}
}

func parseBasicAuthHeader(c *fiber.Ctx) (username string, password string, err error) {
	auth := c.Get(consts.AuthorizationHeaderName)
	if auth == "" {
		return "", "", fiber.NewError(fiber.StatusUnauthorized, errs.EmptyCredentials.Error())
	}

	parts := strings.SplitN(auth, " ", 2)
	if len(parts) != 2 || parts[0] != consts.AuthorizationType {
		return "", "", fiber.NewError(fiber.StatusUnauthorized, errs.EmptyCredentials.Error())
	}

	payload, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", "", fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		return "", "", fiber.NewError(fiber.StatusUnauthorized, errs.EmptyCredentials.Error())
	}

	return pair[0], pair[1], nil
}
