package users

import (
	"fiber-simple-api/domains/sismor/model"
	"fiber-simple-api/repository"
	"log"

	"github.com/gofiber/fiber/v2"
)

type UsersController struct {
	respository repository.IBaseRepository
}

func NewUsersController(respository repository.IBaseRepository) *UsersController {
	return &UsersController{
		respository: respository,
	}
}

func (controller *UsersController) List(ctx *fiber.Ctx) error {
	var err error
	response, err := controller.respository.GetUsers()
	if err != nil {
		ctx.JSON(err)
		return err
	}
	ctx.JSON(response)
	return nil
}

func (controller *UsersController) Detail(ctx *fiber.Ctx) error {
	params := struct {
		ID int64 `params:"id"`
	}{}
	err := ctx.ParamsParser(&params)
	if err != nil {
		log.Println(err)
		ctx.JSON(err)
		return err
	}
	user, err := controller.respository.GetUser(params.ID)
	if err != nil {
		log.Println(err)
		ctx.JSON(err)
		return err
	}
	ctx.JSON(user)
	return nil
}

func (controller *UsersController) Insert(ctx *fiber.Ctx) error {
	userToInsert := new(model.Users)
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
	params := struct {
		ID int64 `params:"id"`
	}{}
	err := ctx.ParamsParser(&params)
	if err != nil {
		log.Println(err)
		ctx.JSON(err)
		return err
	}
	err = controller.respository.DeleteUser(params.ID)
	if err != nil {
		log.Println(err)
		ctx.JSON(err)
		return err
	}
	ctx.Status(fiber.StatusOK)
	return nil
}

func (controller *UsersController) Update(ctx *fiber.Ctx) error {
	params := struct {
		ID int64 `params:"id"`
	}{}
	err := ctx.ParamsParser(&params)
	if err != nil {
		return err
	}

	userData := &model.Users{}
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

	err = controller.respository.UpdateUser(params.ID, userData)
	if err != nil {
		ctx.JSON(err)
		return err
	}
	return nil
}
