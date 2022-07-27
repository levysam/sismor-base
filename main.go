package main

import (
	routes "fiber-simple-api/Routes"

	"github.com/gofiber/fiber/v2"
)

func setupRouter() *fiber.App {
	r := fiber.New()
	routes.Auth(r)
	routes.Health(r)
	routes.Users(r)
	return r
}

func main() {
	r := setupRouter()
	r.Listen(":8080")
}
