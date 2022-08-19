package users

import (
	"fiber-simple-api/database"
	"fiber-simple-api/domains/sismor/model"
	"fiber-simple-api/domains/sismor/table"
	"fiber-simple-api/repository"
	"fmt"

	"github.com/go-jet/jet/v2/mysql"
)

var UsersRepositoryVar repository.IBaseRepository

type UsersRepository struct {
	database *database.Database
}

func init() {
	UsersRepositoryVar = &UsersRepository{}
}

func NewUsersRepository(database *database.Database) *UsersRepository {
	return &UsersRepository{
		database: database,
	}

}

func (repository *UsersRepository) GetUsers() ([]*model.Users, error) {
	var users []*model.Users

	rows, err := repository.database.Query(`select id, name, email from users`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := new(model.Users)
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repository *UsersRepository) GetUser(id int64) (*model.Users, error) {
	user := new(model.Users)
	err := table.Users.SELECT(table.Users.AllColumns).
		WHERE(table.Users.ID.EQ(mysql.Int(id))).
		Query(repository.database, user)
	return user, err
}

func (repository *UsersRepository) GetUserByEmail(email string) (*model.Users, error) {
	user := new(model.Users)
	row := repository.database.QueryRow(`select id, name, email, passwordfrom users where email = ? limit 1`, email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (repository *UsersRepository) InsertUser(user *model.Users) error {
	_, err := table.Users.
		INSERT(table.Users.Name, table.Users.Password, table.Users.Email).
		VALUES(user.Name, user.Password, user.Email).
		Exec(repository.database)
	return err
}

func (repository *UsersRepository) DeleteUser(Id int64) error {
	_, err := table.Users.DELETE().
		WHERE(table.Users.ID.EQ(mysql.Int(Id))).
		Exec(repository.database)
	return err
}

func (repository *UsersRepository) UpdateUser(id int64, UserData *model.Users) error {
	statement := table.Users.UPDATE(table.Users.MutableColumns).
		MODEL(UserData).
		WHERE(table.Users.ID.EQ(mysql.Int(id)))
	fmt.Println(statement.Sql())
	_, err := statement.Exec(repository.database)
	return err
}
