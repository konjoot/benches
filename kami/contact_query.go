package main

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewContactQuery(page int, perPage int) *ContactQuery {
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 1
	}

	return &ContactQuery{
		limit:      perPage,
		offset:     perPage * (page - 1),
		conn:       NewDBConn(),
		collection: NewContactList(),
	}
}

type ContactQuery struct {
	limit      int
	offset     int
	collection []Contact
	conn       *sqlx.DB
}

func (cq *ContactQuery) All() ContactList {
	err := cq.fillUsers()
	if err != nil {
		fmt.Println(err)
		return NewContactList()
	}
	return cq.collection
}

func (cq *ContactQuery) fillUsers() (err error) {
	ps, err := cq.selectUsersStmt()
	if err != nil {
		return
	}
	fmt.Println(ps)

	err = ps.Select(cq.collection, cq.limit, cq.offset)

	return
}

func (cq *ContactQuery) selectUsersStmt() (*sqlx.Stmt, error) {
	if cq.conn == nil {
		return nil, errors.New("Can't connect to DB")
	}

	return cq.conn.Preparex(`
    select  id,
            email,
            first_name,
            last_name,
            middle_name,
            date_of_birth,
            sex
      from users
      where deleted_at is null
      order by id
      limit ?
      offset ?`)
}
