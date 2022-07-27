package routes

import (
	users "fiber-simple-api/Domains/Users"

	"github.com/gofiber/fiber/v2"
)

func Users(r *fiber.App) {
	r.Get("/usuarios", users.List)
	r.Get("/user/:id", users.Detail)
}
