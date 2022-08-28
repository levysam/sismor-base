package repository

import (
	"fiber-simple-api/database"
	"fiber-simple-api/domains/users/repository"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type repositoryFactory struct {
	database *database.Database
}

func NewRepositoryFactory() *repositoryFactory {
	return &repositoryFactory{
		database: nil,
	}
}

func (factory *repositoryFactory) GetRepository(controllerType string) (IBaseRepository, error) {
	factory.GetDatabase()
	if controllerType == "users" {
		return repository.NewUsersRepository(factory.database), nil
	}
	return nil, fmt.Errorf("wrong controller type passed")
}

func (factory *repositoryFactory) GetDatabase() {
	if factory.database != nil {
		return
	}
	godotenv.Load()
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	connection := os.Getenv("DB_CONNECTION")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, connection, port, dbName)
	db, err := database.NewDb(dsn)
	if err != nil {
		return
	}
	factory.database = db
}
