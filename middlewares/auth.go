package middlewares

import (
	localSession "fiber-simple-api/services/sessions"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	_, err := localSession.Get(headers["Authorization"])
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	c.Next()
	return nil
}
