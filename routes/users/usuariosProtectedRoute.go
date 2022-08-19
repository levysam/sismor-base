package users

import (
	"fiber-simple-api/domains/users"

	"github.com/gofiber/fiber/v2"
)

type UsuariosProtectedRoute struct {
	ProtectedRoute
}

func (router *UsuariosProtectedRoute) ListenRoutes(fiber *fiber.App) {
	usersGroup := fiber.Group("/users")
	usersGroup.Get("/", router.ProtectedRoute.Controller.List)
	usersGroup.Get("/:id", router.ProtectedRoute.Controller.Detail)
	usersGroup.Post("/", router.ProtectedRoute.Controller.Insert)
	usersGroup.Patch("/:id", router.ProtectedRoute.Controller.Update)
	usersGroup.Delete("/:id", router.ProtectedRoute.Controller.Delete)
}

func (p *ProtectedRoute) SetController(controller *users.UsersController) {
	p.Controller = controller
}

func (p *ProtectedRoute) GetController() *users.UsersController {
	return p.Controller
}

func (p *ProtectedRoute) SetRepository(repository *users.UsersRepository) {
	p.Repository = repository
}

func (p *ProtectedRoute) GetRepository() *users.UsersRepository {
	return p.Repository
}
