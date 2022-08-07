package users

import (
	"fiber-simple-api/domains/sismor/model"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type repositoryMock struct{}

func (s *repositoryMock) GetUsers() ([]*model.Users, error) {
	var users []*model.Users
	return users, nil
}

func (s *repositoryMock) GetUser(id int64) (*model.Users, error) {
	var user *model.Users
	return user, nil
}
func (s *repositoryMock) GetUserByEmail(email string) (*model.Users, error) {
	var user *model.Users
	return user, nil
}
func (s *repositoryMock) InsertUser(user *model.Users) error {
	return nil
}
func (s *repositoryMock) DeleteUser(Id int64) error {
	return nil
}
func (s *repositoryMock) UpdateUser(id int64, UserData *model.Users) error {
	return nil
}

type repositoryMockError struct{}

type GetUsersError struct{}

func (m *GetUsersError) Error() string {
	return "Erro ao pegar usuários"
}

func (s *repositoryMockError) GetUsers() ([]*model.Users, error) {
	userMock := &model.Users{
		// ID:       1,
		// Name:     "levy",
		// Password: "123",
		// Email:    "levy@123.com",
	}

	return []*model.Users{userMock}, &GetUsersError{}
}

func (s *repositoryMockError) GetUser(id int64) (*model.Users, error) {
	var user *model.Users
	return user, &GetUsersError{}
}
func (s *repositoryMockError) GetUserByEmail(email string) (*model.Users, error) {
	var user *model.Users
	return user, nil
}
func (s *repositoryMockError) InsertUser(user *model.Users) error {
	return nil
}
func (s *repositoryMockError) DeleteUser(Id int64) error {
	return nil
}
func (s *repositoryMockError) UpdateUser(id int64, UserData *model.Users) error {
	return nil
}

func TestUsersController_List(t *testing.T) {
	fiberApp := fiber.New()
	type args struct {
		method string
		url    string
	}
	tests := []struct {
		name    string
		args    *args
		Mock    UsersRepositoryInterface
		wantErr bool
	}{
		{"Success", &args{"get", "/users"}, &repositoryMock{}, false},
		{"Error", &args{"get", "/users"}, &repositoryMockError{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &UsersController{
				respository: tt.Mock,
			}
			request := &fasthttp.RequestCtx{}

			request.Request.Header.SetMethod(tt.args.method)
			request.URI().SetPath(tt.args.url)
			ctxMock := fiberApp.AcquireCtx(request)
			if err := controller.List(ctxMock); (err != nil) != tt.wantErr {
				t.Errorf("UsersController.List() error = %v, wantErr %v", err, tt.wantErr)
			}
			userMock := &model.Users{
				// ID:       1,
				// Name:     "Levy",
				// Password: "123",
				// Email:    "levy@123.com",
			}
			if controller.response != userMock {

			}
		})
	}
}

func TestUsersController_Detail(t *testing.T) {
	fiberApp := fiber.New()
	type args struct {
		method string
		url    string
	}
	tests := []struct {
		name    string
		args    *args
		Mock    UsersRepositoryInterface
		wantErr bool
	}{
		{"Success", &args{"get", "/users/1"}, &repositoryMock{}, false},
		{"Error", &args{"get", "/users/1"}, &repositoryMockError{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &UsersController{
				respository: tt.Mock,
			}
			request := &fasthttp.RequestCtx{}

			request.Request.Header.SetMethod(tt.args.method)
			request.URI().SetPath(tt.args.url)
			ctxMock := fiberApp.AcquireCtx(request)
			if err := controller.Detail(ctxMock); (err != nil) != tt.wantErr {
				t.Errorf("UsersController.List() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsersController_Insert(t *testing.T) {
	type fields struct {
		respository UsersRepositoryInterface
		response    interface{}
	}
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &UsersController{
				respository: tt.fields.respository,
				response:    tt.fields.response,
			}
			if err := controller.Insert(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("UsersController.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
