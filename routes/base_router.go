package routes

import (
	"fiber-simple-api/database"
	"fiber-simple-api/domains/users"
)

type BaseRouter struct{}

func NewBaseRouter() *BaseRouter {
	return &BaseRouter{}
}

func (app *BaseRouter) GetUsersController(Database *database.Database) *users.UsersController {
	repository := users.NewUsersRepository(Database)
	users := users.NewUsersController(repository)

	return users
}
