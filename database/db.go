package database

import (
	"database/sql"
)

type Database struct {
	*sql.DB
}

func NewDb(dsn string) (*Database, error) {
	dbSql, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &Database{
		dbSql,
	}, nil
}
