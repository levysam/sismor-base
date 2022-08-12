package routes

import (
	"fiber-simple-api/database"
	"fiber-simple-api/domains/users"
)

type usuarios struct {
}

func (u *usuarios) MakeRouteProtected(Database *database.Database) iProtectedRoute {
	repository := users.NewUsersRepository(Database)
	return &usuariosProtectedRoute{
		ProtectedRoute: ProtectedRoute{
			repository: repository,
			controller: users.NewUsersController(repository),
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
