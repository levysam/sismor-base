package routes

import (
	"fiber-simple-api/domains/users"

	"github.com/gofiber/fiber/v2"
)

type usuariosProtectedRoute struct {
	ProtectedRoute
}

func (router *usuariosProtectedRoute) ListenRoutes(fiber *fiber.App) {
	usersGroup := fiber.Group("/users")
	usersGroup.Get("/", router.ProtectedRoute.controller.List)
	// usersGroup.Get("/:id", router.Controller.Detail)
	// usersGroup.Post("/", router.Controller.Insert)
	// usersGroup.Patch("/:id", router.Controller.Update)
	// usersGroup.Delete("/:id", router.Controller.Delete)
}

func (p *ProtectedRoute) SetController(controller *users.UsersController) {
	p.controller = controller
}

func (p *ProtectedRoute) GetController() *users.UsersController {
	return p.controller
}

func (p *ProtectedRoute) SetRepository(repository *users.UsersRepository) {
	p.repository = repository
}

func (p *ProtectedRoute) GetRepository() *users.UsersRepository {
	return p.repository
}

func (p *ProtectedRoute) ListenRoutes(app *fiber.App) {
	return
}
