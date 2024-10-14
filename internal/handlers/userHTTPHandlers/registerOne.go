package userHTTPHandlers

import (
	"auth-service-test/internal/domain"
	"auth-service-test/pkg/errs"
	"auth-service-test/pkg/reqvalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// RegisterOneRequest is the request model for registering user
type RegisterOneRequest struct {
	Firstname   string `json:"firstname" validate:"required"`
	Lastname    string `json:"lastname" validate:"required"`
	Email       string `json:"email" validate:"omitempty,email"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,phone"`
	Role        string `json:"role" validate:"required,role"`

	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// RegisterOne godoc
// @Summary Register new user
// @Description Register a new user with all required details
// @Tags users
// @Accept json
// @Produce json
// @Param request body RegisterOneRequest true "New user information"
// @Success 200 {string} string "ok"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Router /register_one [post]
// @Security BasicAuth
func (h *UserHandlers) RegisterOne() fiber.Handler {
	return func(c *fiber.Ctx) error {
		spanName := "UserHandlers.RegisterOne"

		request := RegisterOneRequest{}
		if err := reqvalidator.ReadRequest(c, &request); err != nil {
			return errors.Wrap(errs.InvalidRequest, spanName)
		}

		err := h.userUC.RegisterOne(c.Context(), request.toUserData())
		if err != nil {
			return errors.Wrap(err, spanName)
		}

		return c.JSON(fiber.Map{
			"data": "ok",
		})
	}
}

func (r *RegisterOneRequest) toUserData() domain.User {
	return domain.User{
		Firstname:   r.Firstname,
		Lastname:    r.Lastname,
		Email:       r.Email,
		PhoneNumber: r.PhoneNumber,
		Role:        r.Role,

		Username: domain.Username(r.Username),
		Password: r.Password,
	}
}
