package repository

import "fiber-simple-api/models"

type IBaseRepository interface {
	Find() ([]models.IBaseModel, error)
	GetById(Id int64) (models.IBaseModel, error)
	Insert(data models.IBaseModel) error
	Delete(Id int64) error
	Update(Id int64, data models.IBaseModel) error
}

type IAuthRepository interface {
	IBaseRepository
	GetUserByEmail(email string) (models.IBaseModel, error)
}
