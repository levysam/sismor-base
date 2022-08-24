package users

import (
	"fiber-simple-api/controllers"
	"github.com/gofiber/fiber/v2"
)

type UserRoute struct {
	controller controllers.IBaseController
}

func (router *UserRoute) ListenRoutes(fiber *fiber.App) {
	usersGroup := fiber.Group("/users")
	usersGroup.Get("/", router.controller.List)
	usersGroup.Get("/:id", router.controller.Detail)
	usersGroup.Post("/", router.controller.Insert)
	usersGroup.Patch("/:id", router.controller.Update)
	usersGroup.Delete("/:id", router.controller.Delete)
}
