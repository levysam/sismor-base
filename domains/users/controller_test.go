package users

import (
	"fiber-simple-api/domains/sismor/model"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type repositoryMock struct{}

func (s *repositoryMock) GetUsers() ([]*model.Users, error) {
	userMock := &model.Users{
		ID:       1,
		Name:     "Levy",
		Password: "123",
		Email:    "teste@123.com",
	}

	return []*model.Users{userMock}, nil
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

func (s *repositoryMockError) GetUsers() ([]*model.Users, error) {
	return nil, &GetUsersError{}
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

type GetUsersError struct{}

func (m *GetUsersError) Error() string {
	return "Erro ao pegar usuários"
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
		{"Error", &args{"get", "/users"}, &repositoryMockError{}, true},
		{"Success", &args{"get", "/users"}, &repositoryMock{}, false},
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
			resp := ctxMock.Response().Body()
			var responseError *GetUsersError
			var response []*model.Users

			if tt.wantErr {
				err := fiberApp.Config().JSONDecoder(resp, &responseError)
				if err != nil {
					t.Errorf("Erro ao Decodar o JSON de resposta: %v", err)
				}
				if responseError.Error() != "Erro ao pegar usuários" {
					t.Error("Error on error assert")
				}
				return
			}

			err := fiberApp.Config().JSONDecoder(resp, &response)
			if err != nil {
				t.Errorf("Erro ao Decodar o JSON de resposta: %v", err)
			}

			userMock := &model.Users{
				ID:       1,
				Name:     "Levy",
				Password: "123",
				Email:    "teste@123.com",
			}
			expectedColection := []*model.Users{userMock}
			for user := range response {
				if response[user].ID != expectedColection[user].ID {
					t.Errorf(
						"Json da resposta é diferente do esperado, campo diferente: %v, campo esperado: %v",
						response[user].ID,
						expectedColection[user].ID,
					)
				}
				if response[user].Name != expectedColection[user].Name {
					t.Errorf(
						"Json da resposta é diferente do esperado, campo diferente: %v, campo esperado: %v",
						response[user],
						expectedColection[user],
					)
				}
				if response[user].Password != expectedColection[user].Password {
					t.Errorf(
						"Json da resposta é diferente do esperado, campo diferente: %v, campo esperado: %v",
						response[user],
						expectedColection[user],
					)
				}
				if response[user].Email != expectedColection[user].Email {
					t.Errorf(
						"Json da resposta é diferente do esperado, campo diferente: %v, campo esperado: %v",
						response[user],
						expectedColection[user],
					)
				}
			}
		})
	}
}
