package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
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

func NewDBConn() (db *sqlx.DB) {
	db, err := sqlx.Connect("postgres", DBUrl)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return
}
