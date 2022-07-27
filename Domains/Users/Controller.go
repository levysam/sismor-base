package users

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func List(ctx *fiber.Ctx) error {
	users, err := GetUsers()
	if err != nil {
		log.Println(err)
		ctx.JSON(err)
		return err
	}
	ctx.JSON(users)
	return nil
}

func Detail(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		log.Println(err)
		ctx.JSON(err)
		return err
	}
	user, err := GetUser(id)
	if err != nil {
		log.Println(err)
		ctx.JSON(err)
		return err
	}
	ctx.JSON(user)
	return nil
}
