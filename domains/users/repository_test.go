package users

import (
	"reflect"
	"testing"

	"fiber-simple-api/database"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestNewUsersRepository(t *testing.T) {
	mockDB, _, err := sqlmock.NewWithDSN("testando")
	if err != nil {
		t.Error(err)
	}
	databaseMock := &database.Database{
		mockDB,
	}
	repository := NewUsersRepository(databaseMock)
	if reflect.TypeOf(repository).String() != "*users.UsersRepository" {
		t.Errorf("expecting *users.UsersRepository type but got %v", reflect.TypeOf(repository))
	}
}

// func TestGetUsers(t *testing.T) {
// 	mockDB, mock, err := sqlmock.NewWithDSN("testando")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	RowsMock := sqlmock.NewRows(
// 		[]string{
// 			"id",
// 			"name",
// 			"email",
// 			"password",
// 		},
// 	)

// 	RowsMock.AddRow(
// 		1,
// 		"Diego Sampaio",
// 		"diiegosampaio@gmail.com",
// 		"$2y$10$sqovj3WYQY/z3dZn0pwWG.nN6UOOIR7oWY6EaUY3hdTlYi0TsSoba",
// 	)

// 	QueryMock := mock.ExpectQuery(`select id, name, email from users`)
// 	QueryMock.WillReturnRows(RowsMock)
// 	databaseMock := &database.Database{
// 		mockDB,
// 	}
// 	repository := users.NewUsersRepository(databaseMock)
// 	usersRes, err := repository.GetUsers()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	UserMock := &users.User{
// 		Id:    1,
// 		Name:  "Diego Sampaio",
// 		Email: "diiegosampaio@gmail.com",
// 	}
// 	var usersMock []*users.User
// 	usersMock = append(usersMock, UserMock)
// 	if usersRes == nil {
// 		t.Error("expecting []*users.User and got nil")
// 	}
// 	if usersMock[0].Id != usersRes[0].Id {
// 		t.Error("Mocked user id is different than result")
// 	}
// 	if usersMock[0].Email != usersRes[0].Email {
// 		t.Error("Mocked user email is different than result")
// 	}
// 	if usersMock[0].Name != usersRes[0].Name {
// 		t.Error("Mocked user Name is different than result")
// 	}
// }
