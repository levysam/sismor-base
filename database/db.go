package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func NewDb() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/rodacoop")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
