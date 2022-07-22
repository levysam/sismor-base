package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func NewDb() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/rodacoop")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
