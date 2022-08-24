package routes

import (
	"fiber-simple-api/domains/users"
	"fmt"
)

func GetRoutes(routerType string) (iBaseRouter, error) {
	if routerType == "user" {
		return &users.UserRoute{}, nil
	}
	return nil, fmt.Errorf("wrong controller type passed")
}
