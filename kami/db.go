package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DBHOST   string = "localhost"
	DATABASE string = "lms2_db_dev"
	DBUSER   string = "lms2_db_user"
	DBPASS   string = "lms_2014"
	SSLMODE  string = "disable"
)

var DBUrl string = fmt.Sprintf(
	"postgres://%s:%s@%s/%s?sslmode=%s",
	DBUSER, DBPASS, DBHOST, DATABASE, SSLMODE)

func NewDBConn() (db *sql.DB) {
	db, err := sql.Open("postgres", DBUrl)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return
}
