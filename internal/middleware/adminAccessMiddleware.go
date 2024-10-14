package middleware

import (
	"auth-service-test/internal/domain"
	"auth-service-test/pkg/consts"
	"auth-service-test/pkg/errs"
	"github.com/gofiber/fiber/v2"
)

func (mw *MiddlewareManager) AdminAccess() fiber.Handler {
	return func(c *fiber.Ctx) error {
		username, ok := c.Locals(consts.UsernameFieldName).(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).SendString(errs.EmptyCredentials.Error())
		}

		user, err := mw.userUC.GetOne(c.Context(), domain.Username(username))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}

		if user.Role != consts.AdminRole {
			return c.Status(fiber.StatusForbidden).SendString(errs.NotEnoughRights.Error())
		}

		return c.Next()
	}
}
