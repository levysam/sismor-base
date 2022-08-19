package repository

import "fiber-simple-api/domains/sismor/model"

type IBaseRepository interface {
	GetUsers() ([]*model.Users, error)
	GetUser(id int64) (*model.Users, error)
	GetUserByEmail(email string) (*model.Users, error)
	InsertUser(user *model.Users) error
	DeleteUser(Id int64) error
	UpdateUser(id int64, UserData *model.Users) error
}
