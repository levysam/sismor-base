package controllers

import (
	"fiber-simple-api/domains/users"
	"fiber-simple-api/repository"
	"fmt"
)

func GetController(controllerType string, repository repository.IBaseRepository) (IBaseController, error) {
	if controllerType == "users" {
		return users.NewUsersController(repository), nil
	}
	return nil, fmt.Errorf("wrong controller type passed")
}
