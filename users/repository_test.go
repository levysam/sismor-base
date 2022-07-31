package users_test

import (
	"reflect"
	"testing"

	"fiber-simple-api/database"
	"fiber-simple-api/users"

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
	repository := users.NewUsersRepository(databaseMock)
	if reflect.TypeOf(repository).String() != "*users.UsersRepository" {
		t.Errorf("expecting *users.UsersRepository type but got %v", reflect.TypeOf(repository))
	}
}

func TestGetUsers(t *testing.T) {
	mockDB, mock, err := sqlmock.NewWithDSN("testando")
	if err != nil {
		t.Error(err)
	}

	RowsMock := sqlmock.NewRows(
		[]string{
			"id",
			"name",
			"email",
			"password",
			"created_at",
			"updated_at",
		},
	)

	RowsMock.AddRow(
		1,
		"Diego Sampaio",
		"diiegosampaio@gmail.com",
		"$2y$10$sqovj3WYQY/z3dZn0pwWG.nN6UOOIR7oWY6EaUY3hdTlYi0TsSoba",
		"2021-11-30 18:08:53",
		"2022-02-28 23:07:11",
	)

	QueryMock := mock.ExpectQuery(`select id, name, email, password,created_at, updated_at from rodacoop.users`)
	QueryMock.WillReturnRows(RowsMock)
	databaseMock := &database.Database{
		mockDB,
	}
	repository := users.NewUsersRepository(databaseMock)
	usersRes, err := repository.GetUsers()
	if err != nil {
		t.Error(err)
	}
	UserMock := &users.User{
		Id:         1,
		Name:       "Diego Sampaio",
		Email:      "diiegosampaio@gmail.com",
		Password:   "$2y$10$sqovj3WYQY/z3dZn0pwWG.nN6UOOIR7oWY6EaUY3hdTlYi0TsSoba",
		Created_at: "2021-11-30 18:08:53",
		Updated_at: "2022-02-28 23:07:11",
	}
	var usersMock []*users.User
	usersMock = append(usersMock, UserMock)
	if usersRes == nil {
		t.Error("expecting []*users.User and got nil")
	}
	// t.Error(usersMock[0], usersRes[0])
	if usersMock[0].Id != usersRes[0].Id {
		t.Error("Mocked user id is different than result")
	}
	if usersMock[0].Email != usersRes[0].Email {
		t.Error("Mocked user email is different than result")
	}
	if usersMock[0].Name != usersRes[0].Name {
		t.Error("Mocked user Name is different than result")
	}
	if usersMock[0].Password != usersRes[0].Password {
		t.Error("Mocked user Password is different than result")
	}
	if usersMock[0].Created_at != usersRes[0].Created_at {
		t.Error("Mocked user Created_at is different than result")
	}
	if usersMock[0].Updated_at != usersRes[0].Updated_at {
		t.Error("Mocked user Updated_at is different than result")
	}
}
