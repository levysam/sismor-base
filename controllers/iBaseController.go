package controllers

import "github.com/gofiber/fiber/v2"

type IBaseController interface {
	List(ctx *fiber.Ctx) error
	Detail(ctx *fiber.Ctx) error
	Insert(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}
