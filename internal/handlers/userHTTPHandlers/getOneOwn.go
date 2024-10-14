package userHTTPHandlers

import (
	"auth-service-test/internal/domain"
	"auth-service-test/pkg/consts"
	"auth-service-test/pkg/errs"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// GetOneOwn godoc
// @Summary Get own user information
// @Description Get details of the currently authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} GetOneResponse
// @Failure 401 {string} string "Unauthorized"
// @Router /get_one_own [get]
// @Security BasicAuth
func (h *UserHandlers) GetOneOwn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		spanName := "UserHandlers.GetOneOwn"

		username, ok := c.Locals(consts.UsernameFieldName).(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).SendString(errs.EmptyCredentials.Error())
		}

		user, err := h.userUC.GetOne(c.Context(), domain.Username(username))
		if err != nil {
			return errors.Wrap(err, spanName)
		}

		return c.JSON(fiber.Map{
			"data": toGetOneResponse(user),
		})
	}
}
