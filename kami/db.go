package main

import (
	"github.com/jackc/pgx"
)

var connConfig = pgx.ConnConfig{
	Host:      "localhost",
	Database:  "lms2_db_dev",
	User:      "lms2_db_user",
	Password:  "lms_2014",
	TLSConfig: nil,
}

var config = pgx.ConnPoolConfig{
	ConnConfig:     connConfig,
	MaxConnections: 5,
}

var dbPool *pgx.ConnPool

func DBConn() (*pgx.ConnPool, error) {
	if dbPool != nil {
		return dbPool, nil
	}

	var err error

	dbPool, err = pgx.NewConnPool(config)
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}
