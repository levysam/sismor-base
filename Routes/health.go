package routes

import "github.com/gofiber/fiber/v2"

func Health(ctx *fiber.App) {
	ctx.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}
