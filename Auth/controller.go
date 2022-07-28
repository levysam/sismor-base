package auth

import (
	"fiber-simple-api/services"
	localSession "fiber-simple-api/services/sessions"
	"fiber-simple-api/users"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	loginForm := new(LoginForm)
	err := c.BodyParser(loginForm)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}
	user, err := users.GetUserByEmail(loginForm.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Nenhum Usuário encontrado")
	}

	if loginForm.Password != user.Password {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token, err := services.CreateJwtToken(user.Id)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	localSession.Set(token.AccessToken, []byte("1"), token.AtExpires)
	c.Status(fiber.StatusAccepted).JSON(fiber.Map{"AccessToken": token.AccessToken, "RefreshToken": token.RefreshToken})
	return nil
}
