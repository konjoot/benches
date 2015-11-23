package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DBHOST   string = "localhost"
	DATABASE string = "lms2_development_2"
	DBUSER   string = "lms"
	DBPASS   string = ""
	SSLMODE  string = "disable"
)

var DBUrl string = fmt.Sprintf(
	"postgres://%s:%s@%s/%s?sslmode=%s",
	DBUSER, DBPASS, DBHOST, DATABASE, SSLMODE)

var db *sql.DB

func DBConn() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	var err error

	db, err = sql.Open("postgres", DBUrl)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(75)

	return db, nil
}
