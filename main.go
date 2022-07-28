package main

import (
	"fiber-simple-api/routes"
	"log"

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
	log.Fatal(r.Listen(":8080"))
}
