package routes

import (
	authController "fiber-simple-api/Auth"
	"fiber-simple-api/database"
)

type authProtectedRoute struct {
	ProtectedRoute
}

func (a *authProtectedRoute) MakeRouteProtected(Database *database.Database) iProtectedRoute {
	return &authProtectedRoute{
		ProtectedRoute: ProtectedRoute{
			controller: authController.NewAuthController(),
		},
	}
}
