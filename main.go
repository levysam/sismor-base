package main

import (
	"fiber-simple-api/database"
	"fiber-simple-api/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type BaseApp struct {
	Database  *database.Database
	FiberBase *fiber.App
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
	app := fiber.New()
	if err != nil {
		log.Fatal(err)
	}
	return &BaseApp{
		Database:  db,
		FiberBase: app,
	}
}

func main() {
	app := NewApp()
	//app.RouterBase.GetUsersController(app.Database)
	usersFactory, err := routes.GetRoutes("user")
	if err != nil {
		log.Fatal(err)
	}
	usersFactory.ListenRoutes(app.FiberBase)
	//printShoeDetails(nikeShoe)
	log.Fatal(app.FiberBase.Listen(":8080"))
}
