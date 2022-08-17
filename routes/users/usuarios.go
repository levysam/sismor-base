package users

import (
	"fiber-simple-api/database"
	"fiber-simple-api/domains/users"
)

type Usuarios struct {
}

func (u *Usuarios) MakeRouteProtected(Database *database.Database) IProtectedRoute {
	repository := users.NewUsersRepository(Database)
	return &UsuariosProtectedRoute{
		ProtectedRoute: ProtectedRoute{
			Repository: repository,
			Controller: users.NewUsersController(repository),
		},
	}
}

// func (u *usuarios) MakeRoutePublic() iShort {
// 	return &adidasShort{
// 		short: short{
// 			logo: "adidas",
// 			size: 14,
// 		},
// 	}
// }
