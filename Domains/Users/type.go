package users

import (
	_ "github.com/go-sql-driver/mysql"
	database "github.com/levysam/sismor-base/database"
)

type Users struct {
	id       int64
	name     string
	email    string
	password string
	isAdmin  bool
}

func GetUsers() []Users {
	db := database.NewDb()
	var users []Users
	db.Select(users, "select id, name, email, password, isAdmin from rodacoop.users limit 100")
	return users
}
