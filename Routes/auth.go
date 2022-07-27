package routes

import (
	auth "fiber-simple-api/Domains/Auth"

	"github.com/gofiber/fiber/v2"
)

func Auth(router *fiber.App) {
	router.Post("/login", auth.Login)
}
