package main

import (
	"fiber-simple-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type BaseApp struct {
	FiberBase *fiber.App
}

func NewApp() *BaseApp {
	godotenv.Load()
	app := fiber.New()
	return &BaseApp{
		FiberBase: app,
	}
}

func main() {
	app := NewApp()
	//app.RouterBase.GetUsersController(app.Database)
	usersFactory, err := routes.GetRoutes("users")
	if err != nil {
		log.Fatal(err)
	}
	usersFactory.ListenRoutes(app.FiberBase)
	//printShoeDetails(nikeShoe)
	log.Fatal(app.FiberBase.Listen(":8080"))
}
