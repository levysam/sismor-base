package routes

import (
	"fiber-simple-api/controllers"
	"fiber-simple-api/domains/users"
	"fmt"
	"log"
)

func GetRoutes(routerType string) (iBaseRouter, error) {
	if routerType == "users" {
		controller, err := controllers.GetController(routerType)
		if err != nil {
			log.Fatal(err)
		}
		userRoute := users.NewUserRoute(controller)
		return userRoute, nil
	}
	return nil, fmt.Errorf("wrong controller type passed")
}
