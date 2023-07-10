package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	ddb, err := sql.Open("postgres", "dbname=arunraghunath sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	DB = ddb
	return
}
