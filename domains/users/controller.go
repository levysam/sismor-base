package users

import (
	"fiber-simple-api/domains/sismor/model"
	"log"

	"github.com/gofiber/fiber/v2"
)

var UsersControllerVar usersControllerInterface

func init() {
	UsersControllerVar = &UsersController{}
}

type UsersController struct {
	repository UsersRepositoryInterface
}

type usersControllerInterface interface {
	List(ctx *fiber.Ctx) error
	Detail(ctx *fiber.Ctx) error
	Insert(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

func NewUsersController(repository UsersRepositoryInterface) *UsersController {
	return &UsersController{
		repository: repository,
	}
}

func (controller *UsersController) List(ctx *fiber.Ctx) error {
	var err error
	response, err := controller.repository.GetUsers()
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
	user, err := controller.repository.GetUser(params.ID)
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
	err = controller.repository.InsertUser(userToInsert)
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
	err = controller.repository.DeleteUser(params.ID)
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
	// oldUser, err := controller.repository.GetUser(id)
	if err != nil {
		ctx.JSON(err)
		return err
	}
	err = ctx.BodyParser(userData)
	if err != nil {
		ctx.JSON(err)
		return err
	}

	err = controller.repository.UpdateUser(params.ID, userData)
	if err != nil {
		ctx.JSON(err)
		return err
	}
	return nil
}
