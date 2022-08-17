package auth

import (
	"fiber-simple-api/domains/users"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var AuthControllerVar authControllerInterface

func init() {
	AuthControllerVar = &AuthController{}
}

type AuthController struct {
	respository users.UsersRepositoryInterface
	response    interface{}
}

type authControllerInterface interface {
	Login(c *fiber.Ctx, repository *users.UsersRepository) error
}

func NewAuthController(respository authControllerInterface) *AuthController {
	return &AuthController{
		respository: respository,
	}
}

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

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
