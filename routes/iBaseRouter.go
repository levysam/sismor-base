package routes

import "fiber-simple-api/database"

type iBaseRouter interface {
	MakeRouteProtected(Database *database.Database) iProtectedRoute
	//MakeRoutePublic() iPublicRoute
}

type BaseRouterV2 struct {
	route string
}

func NewBaseRouterV2(name string) *BaseRouterV2 {
	return &BaseRouterV2{route: name}
}

func (c BaseRouterV2) Route() iBaseRouter {
	if c.route == "auth" {
		return &auth{}
	}
	if c.route == "usuarios" {
		return &usuarios{}
	}
	// if c.route == "health" {
	// 	return &nike{}
	// }
	return nil
}
