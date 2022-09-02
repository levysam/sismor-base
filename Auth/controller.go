package auth

import (
	"fiber-simple-api/models"
	"fiber-simple-api/repository"
	"log"

	"github.com/gofiber/fiber/v2"
)

type IAuthController interface {
	Login(c *fiber.Ctx) error
}

type IAuthRepository interface {
	repository.IBaseRepository
	GetUserByEmail(email string) (models.IBaseModel, error)
}

type AuthController struct {
	repository IAuthRepository
}

func NewAuthController() IAuthController {
	repository := repository.NewRepositoryFactory()
	userRepository, err := repository.GetRepositoryAuth()
	if err != nil {
		log.Fatal(err)
	}
	return &AuthController{
		repository: userRepository,
	}
}

func (Controller AuthController) Login(c *fiber.Ctx) error {

	loginForm := new(LoginForm)
	err := c.BodyParser(loginForm)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}
	user, err := Controller.repository.GetUserByEmail(loginForm.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Nenhum Usuário encontrado")
	}

	if loginForm.Password != user.Password {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return nil
}
