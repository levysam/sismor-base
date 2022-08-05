package users

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int64  `json:",omitempty"`
	Name     string `json:",omitempty"`
	Email    string `json:",omitempty"`
	Password string `json:",omitempty"`
}
