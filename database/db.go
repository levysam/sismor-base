package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	*sql.DB
}

var Open = sql.Open

func NewDb(dsn string) (*Database, error) {
	dbSql, err := Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &Database{
		dbSql,
	}, nil
}
