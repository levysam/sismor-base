package users

import (
	"fiber-simple-api/domains/users"

	"github.com/gofiber/fiber/v2"
)

type IProtectedRoute interface {
	SetController(controller *users.UsersController)
	SetRepository(repository *users.UsersRepository)
	GetController() *users.UsersController
	GetRepository() *users.UsersRepository
	ListenRoutes(*fiber.App)
}

type ProtectedRoute struct {
	Controller *users.UsersController
	Repository *users.UsersRepository
}
