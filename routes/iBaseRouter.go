package routes

import "github.com/gofiber/fiber/v2"

type iBaseRouter interface {
	ListenRoutes(fiber *fiber.App)
}
