package routes

import (
	auth "fiber-simple-api/Auth"
	"fiber-simple-api/domains/users"

	"github.com/gofiber/fiber/v2"
)

type authsProtectedRoute struct {
	ProtectedRoute
}

func (router *authsProtectedRoute) ListenRoutes(fiber *fiber.App) {
	authGroup := fiber.Group("/login")
	authGroup.Post("/", login)
}

func (p *ProtectedRoute) SetController(controller *auth.AuthController) {
	p.controller = controller
}

func (p *ProtectedRoute) GetController() *auth.AuthController {
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
