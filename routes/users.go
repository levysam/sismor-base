package routes

import (
	"fiber-simple-api/users"
	"fmt"
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
	fmt.Println("passando por essa rota")
	usersGroup := router.Base.Fiber.Group("/users")
	usersGroup.Get("/", router.Controller.List)
	usersGroup.Get("/:id", router.Controller.Detail)
	usersGroup.Post("/", router.Controller.Insert)
	usersGroup.Patch("/:id", router.Controller.Update)
	usersGroup.Delete("/:id", router.Controller.Delete)
}
