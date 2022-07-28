package routes

import (
	"fiber-simple-api/middlewares"
	"fiber-simple-api/users"

	"github.com/gofiber/fiber/v2"
)

func Users(r *fiber.App) {
	usersGroup := r.Group("/users")
	usersGroup.Get("/", middlewares.AuthMiddleware, users.List)
	usersGroup.Get("/:id", middlewares.AuthMiddleware, users.Detail)
}
