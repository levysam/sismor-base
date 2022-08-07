package auth

import (
	"fiber-simple-api/domains/users"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx, repository *users.UsersRepository) error {
	loginForm := new(LoginForm)
	err := c.BodyParser(loginForm)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}
	user, err := repository.GetUserByEmail(loginForm.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Nenhum Usuário encontrado")
	}

	if loginForm.Password != user.Password {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return nil
}
