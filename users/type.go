package users

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int64
	Name     string
	Email    string
	Password string
}
