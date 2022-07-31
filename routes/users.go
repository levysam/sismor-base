package routes

import (
	"fiber-simple-api/users"
)

type UsersRouter struct {
	Controller *users.UsersController
	Base       *BaseRouter
}

func NewUsersRouter(baseRouter *BaseRouter, controller *users.UsersController) *BaseRouter {
	router := &UsersRouter{
		Base:       baseRouter,
		Controller: controller,
	}
	router.Users()
	return router.Base
}

func (router UsersRouter) Users() {
	usersGroup := router.Base.Fiber.Group("/users")
	usersGroup.Get("/", router.Controller.List)
	usersGroup.Get("/:id", router.Controller.Detail)
	usersGroup.Post("/add", router.Controller.Insert)
}
