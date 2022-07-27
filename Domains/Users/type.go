package users

import (
	database "fiber-simple-api/database"

	_ "github.com/go-sql-driver/mysql"
)

var db = database.NewDb()

type User struct {
	Id         int64
	Name       string
	Email      string
	Password   string
	Created_at string
	Updated_at string
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}

func GetUsers() ([]User, error) {
	users := []User{}
	err := db.Select(&users, `select id, name, email, password,created_at, updated_at from rodacoop.users`)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id int) (User, error) {
	user := User{}
	row := db.QueryRow(`select id, name, email, password, created_at, updated_at from rodacoop.users where id = ? limit 1`, id)
	row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at)
	return user, nil
}

func GetUserByEmail(email string) (User, error) {
	user := User{}
	row := db.QueryRow(`select id, name, email, password, created_at, updated_at from rodacoop.users where email = ? limit 1`, email)
	row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at)
	return user, nil
}
