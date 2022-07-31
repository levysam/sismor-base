package users

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type UsersController struct {
	respository *UsersRepository
}

func NewUsersController(respository *UsersRepository) *UsersController {
	return &UsersController{
		respository: respository,
	}
}

func (controller UsersController) List(ctx *fiber.Ctx) error {
	users, err := controller.respository.GetUsers()
	if err != nil {
		log.Println(err)
		ctx.JSON(err)
		return err
	}
	ctx.JSON(users)
	return nil
}

func (controller UsersController) Detail(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		log.Println(err)
		ctx.JSON(err)
		return err
	}
	user, err := controller.respository.GetUser(id)
	if err != nil {
		log.Println(err)
		ctx.JSON(err)
		return err
	}
	ctx.JSON(user)
	return nil
}