package routes

import (
	"github.com/gofiber/fiber/v2"
)

type BaseRouter struct {
	Fiber *fiber.App
}

func NewBaseRouter(fiber *fiber.App) *BaseRouter {
	return &BaseRouter{
		Fiber: fiber,
	}
}
