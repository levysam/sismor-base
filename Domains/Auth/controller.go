package auth

import (
	"log"

	users "fiber-simple-api/Domains/Users"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Login(c *fiber.Ctx) error {
	loginForm := new(LoginForm)
	if err := c.BodyParser(loginForm); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err)
	}
	user, err := users.GetUserByEmail(loginForm.Email)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	// Throws Unauthorized error
	if loginForm.Password != user.Password {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := UserClaims{
		user.Id,
		user.Name,
		user.Email,
		user.Created_at,
		user.Updated_at,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
			Id:        "testing",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	t, err := token.SignedString([]byte("tknjdpaimxalsmdrk"))
	if err != nil {
		log.Printf("O erro é esse: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
