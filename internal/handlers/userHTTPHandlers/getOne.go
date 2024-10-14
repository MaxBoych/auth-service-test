package userHTTPHandlers

import (
	"auth-service-test/internal/domain"
	"auth-service-test/pkg/errs"
	"auth-service-test/pkg/reqvalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// GetOneRequest is the request model for fetching a user by username
type GetOneRequest struct {
	Username string `json:"username" validate:"required"`
}

// GetOneResponse is the response model for a user details
type GetOneResponse struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`

	Username string `json:"username"`
}

// GetOne godoc
// @Summary Get user by username
// @Description Get user details by providing a username
// @Tags users
// @Accept json
// @Produce json
// @Param request body GetOneRequest true "User information"
// @Success 200 {object} GetOneResponse
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Router /get_one [post]
// @Security BasicAuth
func (h *UserHandlers) GetOne() fiber.Handler {
	return func(c *fiber.Ctx) error {
		spanName := "UserHandlers.GetOne"

		request := GetOneRequest{}
		if err := reqvalidator.ReadRequest(c, &request); err != nil {
			return errors.Wrap(errs.InvalidRequest, spanName)
		}

		user, err := h.userUC.GetOne(c.Context(), domain.Username(request.Username))
		if err != nil {
			return errors.Wrap(err, spanName)
		}

		return c.JSON(fiber.Map{
			"data": toGetOneResponse(user),
		})
	}
}

func toGetOneResponse(user domain.User) GetOneResponse {
	return GetOneResponse{
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,

		Username: string(user.Username),
	}
}
