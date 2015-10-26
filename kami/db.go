package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	DBHOST   string = "localhost"
	DATABASE string = "lms2_development"
	DBUSER   string = "lms"
	DBPASS   string = ""
	SSLMODE  string = "disable"
)

var DBUrl string = fmt.Sprintf(
	"postgres://%s:%s@%s/%s?sslmode=%s",
	DBUSER, DBPASS, DBHOST, DATABASE, SSLMODE)

func NewDBConn() (db *sql.DB) {
	db, err := sql.Open("postgres", DBUrl)
	if err != nil {
		log.Print(err)
		return nil
	}

	return
}
