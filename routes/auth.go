package routes

import (
	auth "fiber-simple-api/Auth"

	"github.com/gofiber/fiber/v2"
)

func Auth(router *fiber.App) {
	router.Post("/login", auth.Login)
}
