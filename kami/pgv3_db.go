package main

import (
	"errors"
	"gopkg.in/pg.v3"
)

var pgConfig = &pg.Options{
	Host:     "localhost",
	Database: "lms2_development_2",
	User:     "lms",
	Password: "",
	SSL:      false,
	PoolSize: 50,
}

var pgDB *pg.DB

func PgDBConn() (*pg.DB, error) {
	if pgDB != nil {
		return pgDB, nil
	}

	pgDB = pg.Connect(pgConfig)
	if pgDB == nil {
		return nil, errors.New("Can't connect to DB")
	}

	return pgDB, nil
}
