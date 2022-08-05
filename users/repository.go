package users

import (
	"encoding/json"
	"fiber-simple-api/database"
	model "fiber-simple-api/users/type"
	"fmt"
	"strings"
)

type UsersRepository struct {
	database *database.Database
}

func NewUsersRepository(database *database.Database) *UsersRepository {
	return &UsersRepository{
		database: database,
	}
}

func (repository *UsersRepository) GetUsers() ([]*User, error) {
	var users []*User

	rows, err := repository.database.Query(`select id, name, email from users`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := new(User)
		err = rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repository *UsersRepository) GetUser(id int) (*User, error) {
	user := new(User)
	row := repository.database.QueryRow(`select id, name, email from users where id = ? limit 1`, id)
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return user, err
	}
	return user, err
}

func (repository *UsersRepository) GetUserByEmail(email string) (*User, error) {
	user := new(User)
	row := repository.database.QueryRow(`select id, name, email, passwordfrom users where email = ? limit 1`, email)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (repository *UsersRepository) InsertUser(user *model.Users) error {
	fmt.Println(Users.INSERT(Users.ID, Users.Name).VALUES(user).Exec(repository.database))
	return nil
	// tx, err := repository.database.Begin()
	// if err != nil {
	// 	return err
	// }
	// _, err = tx.Exec(
	// 	"insert into users (name, email, password) values (?, ?, ?)",
	// 	user.Name,
	// 	user.Email,
	// 	user.Password,
	// )
	// if err != nil {
	// 	return err
	// }
	// err = tx.Commit()
	// if err != nil {
	// 	return err
	// }
	// return nil
}

func (repository *UsersRepository) DeleteUser(Id int) error {
	tx, err := repository.database.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		"delete from users where id = ?",
		Id,
	)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repository *UsersRepository) UpdateUser(id int, UserData *User) error {
	// Users.INSERT(Users.ID, Users.Name).VALUES(UserData)
	tx, err := repository.database.Begin()
	if err != nil {
		return err
	}
	statement, data := createQuery(*UserData)
	data = append(data, id)
	_, err = tx.Exec(
		"UPDATE users SET "+statement+" WHERE id=?",
		data...,
	)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func createQuery(UserData interface{}) (string, []interface{}) {
	data := []interface{}{}
	JsonData, _ := json.Marshal(UserData)
	var MapUser map[string]interface{}
	json.Unmarshal(JsonData, &MapUser)
	StatmentString := []string{}
	for key, value := range MapUser {
		StatmentString = append(StatmentString, key+" = ?")
		data = append(data, value)
	}
	return strings.Join(StatmentString, ", "), data
}
