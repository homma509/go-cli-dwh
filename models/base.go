package models

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
)

var DB *sql.DB

func init() {
	var err error
	dsn := "host=127.0.0.1 port=5432 user=test password=honma0 dbname=dvdrental sslmode=disable";
	DB, err = sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatalln(err)
	}
}