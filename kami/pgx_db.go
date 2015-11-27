package main

import "github.com/jackc/pgx"

var connConfig = pgx.ConnConfig{
	Host:      "localhost",
	Database:  "lms2_development_2",
	User:      "lms",
	Password:  "",
	TLSConfig: nil,
}

var pgxConfig = pgx.ConnPoolConfig{
	ConnConfig:     connConfig,
	MaxConnections: 75,
}

var pgxDB *pgx.ConnPool

func PgxDBConn() (*pgx.ConnPool, error) {
	if pgxDB != nil {
		return pgxDB, nil
	}

	var err error

	pgxDB, err = pgx.NewConnPool(pgxConfig)
	if err != nil {
		return nil, err
	}

	return pgxDB, nil
}
