package database_test

import (
	"database/sql"
	"fiber-simple-api/database"
	"reflect"
	"testing"
)

type DbError struct{}

func TestNewDB(t *testing.T) {
	database.Open = OpenMock
	db, err := database.NewDb("teste")
	if err != nil {
		t.Errorf("Erro ao criar banco")
	}
	if reflect.TypeOf(db).String() != "*database.Database" {
		t.Error("Tipo retornado incorreto")
	}
}

func TestNewDBError(t *testing.T) {
	database.Open = OpenMock
	db, err := database.NewDb("error")
	if err == nil {
		t.Error("Banco criado mesmo com dsn incorreto")
	}
	if db != nil {
		t.Error("Tipo retornado incorreto")
	}
}

func OpenMock(driver string, dsn string) (*sql.DB, error) {
	if dsn == "error" {
		return nil, &DbError{}
	}
	return &sql.DB{}, nil
}

func (err *DbError) Error() string {
	return "Banco não criado"
}
