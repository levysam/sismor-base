package routes

import (
	"fiber-simple-api/database"
	"fiber-simple-api/routes/users"
)

type iBaseRouter interface {
	MakeRouteProtected(Database *database.Database) users.IProtectedRoute
	//MakeRoutePublic() iPublicRoute
}

type BaseRouter struct {
	route string
}

func NewBaseRouter(name string) *BaseRouter {
	return &BaseRouter{route: name}
}

func (c BaseRouter) Route() iBaseRouter {
	if c.route == "usuarios" {
		return &users.Usuarios{}
	}
	// if c.route == "health" {
	// 	return &nike{}
	// }
	return nil
}
