package users

import (
	_ "github.com/go-sql-driver/mysql"
	database "github.com/levysam/sismor-base/database"
)

var db = database.NewDb()

type User struct {
	Id         int64
	Name       string
	Email      string
	Created_at string
	Updated_at string
}

func GetUsers() ([]User, error) {
	users := []User{}
	err := db.Select(&users, `select id, name, email, created_at, updated_at from rodacoop.users`)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id int64) (User, error) {
	user := User{}
	row := db.QueryRow(`select id, name, email, created_at, updated_at from rodacoop.users where id = ? limit 1`, id)
	row.Scan(&user.Id, &user.Name, &user.Email, &user.Created_at, &user.Updated_at)
	return user, nil
}
