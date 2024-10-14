package userHTTPHandlers

import (
	"auth-service-test/internal/domain"
	"auth-service-test/pkg/consts"
	"auth-service-test/pkg/errs"
	"auth-service-test/pkg/reqvalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// DeleteOneRequest is the request model for deleting user by username
type DeleteOneRequest struct {
	Username string `json:"username" validate:"required"`
}

// DeleteOne godoc
// @Summary Delete user
// @Description Delete a user by username
// @Tags users
// @Accept json
// @Produce json
// @Param request body DeleteOneRequest true "User delete information"
// @Success 200 {string} string "ok"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Router /delete_one [post]
// @Security BasicAuth
func (h *UserHandlers) DeleteOne() fiber.Handler {
	return func(c *fiber.Ctx) error {
		spanName := "UserHandlers.DeleteOne"

		request := DeleteOneRequest{}
		if err := reqvalidator.ReadRequest(c, &request); err != nil {
			return errors.Wrap(errs.InvalidRequest, spanName)
		}

		username, ok := c.Locals(consts.UsernameFieldName).(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).SendString(errs.EmptyCredentials.Error())
		}

		if request.Username == username {
			return c.Status(fiber.StatusForbidden).SendString(errs.CannotDeleteYourself.Error())
		}

		err := h.userUC.DeleteOne(c.Context(), domain.Username(request.Username))
		if err != nil {
			return errors.Wrap(err, spanName)
		}

		return c.JSON(fiber.Map{
			"data": "ok",
		})
	}
}
