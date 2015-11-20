package main

import (
	// "database/sql"
	// "fmt"

	// "gopkg.in/pg.v3"
	// _ "github.com/lib/pq"
	"github.com/jackc/pgx"
)

// const (
// 	DBHOST   string = "localhost"
// 	DATABASE string = "lms2_development_2"
// 	DBUSER   string = "lms"
// 	DBPASS   string = ""
// 	SSLMODE  string = "disable"
// )

var connConfig = pgx.ConnConfig{
	Host:     "localhost",
	Database: "lms2_development_2",
	User:     "lms",
	Password: "",
}

var opts = pgx.ConnPoolConfig{
	ConnConfig:     connConfig,
	MaxConnections: 10,
}

// var DBUrl string = fmt.Sprintf(
// 	"postgres://%s:%s@%s/%s?sslmode=%s",
// 	DBUSER, DBPASS, DBHOST, DATABASE, SSLMODE)

var db *pgx.ConnPool

func DbPool() (*pgx.ConnPool, error) {
	if db != nil {
		return db, nil
	}

	var err error

	// db, err = sql.Open("pgx", DBUrl)
	db, err = pgx.NewConnPool(opts)
	if err != nil {
		return nil, err
	}

	// db.SetMaxIdleConns(75)
	// db.SetMaxOpenConns(90)

	return db, nil
}

func NewConn() (conn *pgx.Conn, err error) {
	db, err := DbPool()
	if err != nil {
		return
	}

	conn, err = db.Acquire()
	if err != nil {
		return
	}

	return
}
