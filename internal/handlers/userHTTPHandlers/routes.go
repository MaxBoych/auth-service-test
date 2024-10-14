package userHTTPHandlers

import (
	"auth-service-test/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapUserRoutes(group fiber.Router, h Handlers, mw *middleware.MiddlewareManager) {
	group.Get("/get_one_own", mw.BasicAuth(), h.GetOneOwn())
	group.Post("/get_one", mw.BasicAuth(), mw.AdminAccess(), h.GetOne())
	group.Post("/register_one", mw.BasicAuth(), mw.AdminAccess(), h.RegisterOne())
	group.Patch("/change_one", mw.BasicAuth(), mw.AdminAccess(), h.ChangeOne())
	group.Delete("/delete_one", mw.BasicAuth(), mw.AdminAccess(), h.DeleteOne())
}
