package userHTTPHandlers

import (
	"auth-service-test/internal/domain"
	"auth-service-test/pkg/errs"
	"auth-service-test/pkg/reqvalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// ChangeOneRequest is the request model for changing user data by username
type ChangeOneRequest struct {
	Username string `json:"username" validate:"required"`

	Firstname   *string `json:"firstname" validate:"omitempty"`
	Lastname    *string `json:"lastname" validate:"omitempty"`
	Email       *string `json:"email" validate:"omitempty,email"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,phone"`
	Role        *string `json:"role" validate:"omitempty,role"`
}

// ChangeOne godoc
// @Summary Update user details
// @Description Update details for a specific user by username
// @Tags users
// @Accept json
// @Produce json
// @Param request body ChangeOneRequest true "User update information"
// @Success 200 {string} string "ok"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Router /change_one [post]
// @Security BasicAuth
func (h *UserHandlers) ChangeOne() fiber.Handler {
	return func(c *fiber.Ctx) error {
		spanName := "UserHandlers.ChangeOne"

		request := ChangeOneRequest{}
		if err := reqvalidator.ReadRequest(c, &request); err != nil {
			return errors.Wrap(errs.InvalidRequest, spanName)
		}

		err := h.userUC.ChangeOne(c.Context(), request.toUpdatedUserData())
		if err != nil {
			return errors.Wrap(err, spanName)
		}

		return c.JSON(fiber.Map{
			"data": "ok",
		})
	}
}

func (r *ChangeOneRequest) toUpdatedUserData() domain.UpdatedUserData {
	return domain.UpdatedUserData{
		Username: domain.Username(r.Username),
		
		Firstname:   r.Firstname,
		Lastname:    r.Lastname,
		Email:       r.Email,
		PhoneNumber: r.PhoneNumber,
		Role:        r.Role,
	}
}
