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

func (controller *UsersController) List(ctx *fiber.Ctx) error {
	users, err := controller.respository.GetUsers()
	if err != nil {
		log.Println(err)
		ctx.JSON(err)
		return err
	}
	ctx.JSON(users)
	return nil
}

func (controller *UsersController) Detail(ctx *fiber.Ctx) error {
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

func (controller *UsersController) Insert(ctx *fiber.Ctx) error {
	userToInsert := new(User)
	err := ctx.BodyParser(userToInsert)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}
	err = controller.respository.InsertUser(userToInsert)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}
	return nil
}

func (controller *UsersController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		log.Println(err)
		ctx.JSON(err)
		return err
	}
	err = controller.respository.DeleteUser(id)
	if err != nil {
		log.Println(err)
		ctx.JSON(err)
		return err
	}
	ctx.Status(fiber.StatusOK)
	return nil
}

func (controller *UsersController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	userData := new(User)
	// oldUser, err := controller.respository.GetUser(id)
	if err != nil {
		ctx.JSON(err)
		return err
	}
	err = ctx.BodyParser(userData)
	if err != nil {
		ctx.JSON(err)
		return err
	}

	err = controller.respository.UpdateUser(id, userData)
	if err != nil {
		ctx.JSON(err)
		return err
	}
	return nil
}
