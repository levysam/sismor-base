package main

import (
	"fiber-simple-api/database"
	"fiber-simple-api/routes"
	"fiber-simple-api/users"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type BaseApp struct {
	Database *database.Database
}

func NewApp() *BaseApp {
	godotenv.Load()
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	connection := os.Getenv("DB_CONNECTION")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, connection, port, dbName)
	db, err := database.NewDb(dsn)
	if err != nil {
		log.Fatal(err)
	}
	return &BaseApp{
		Database: db,
	}
}

func (app BaseApp) usersRouter(fiber *fiber.App, controller *users.UsersController) *routes.BaseRouter {
	router := routes.NewBaseRouter(fiber)
	return routes.NewUsersRouter(router, controller)
}

func (app BaseApp) getUsersController() *users.UsersController {
	repository := users.NewUsersRepository(app.Database)
	return users.NewUsersController(repository)
}

func main() {
	fiber := fiber.New()
	app := NewApp()
	usersController := app.getUsersController()
	app.usersRouter(fiber, usersController)
	log.Fatal(fiber.Listen(":3000"))
}
