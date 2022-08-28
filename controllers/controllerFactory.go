package controllers

import (
	"fiber-simple-api/domains/users/controllers"
	"fiber-simple-api/repository"
	"fmt"
	"log"
)

func GetController(controllerType string) (IBaseController, error) {
	repositoryFactory := repository.NewRepositoryFactory()
	repository, err := repositoryFactory.GetRepository(controllerType)
	if err != nil {
		log.Fatal(err)
	}
	if controllerType == "users" {
		return controllers.NewUsersController(repository), nil
	}
	return nil, fmt.Errorf("wrong controller type passed")
}
