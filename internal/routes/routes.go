package routes

package routes

import (
	"github.com/gofiber/fiber/v2"

	"go-user-api/internal/handler"
)

func RegisterRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	api := app.Group("/users")

	api.Get("/", userHandler.ListUsers)
	api.Get("/:id", userHandler.GetUser)
	api.Post("/", userHandler.CreateUser)
	api.Put("/:id", userHandler.UpdateUser)
	api.Delete("/:id", userHandler.DeleteUser)
}
