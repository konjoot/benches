package main

import (
	"database/sql"
	"errors"
	"log"
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
		collection: NewContactList(perPage),
	}
}

type ContactQuery struct {
	limit      int
	offset     int
	collection []*Contact
	conn       *sql.DB
}

func (cq *ContactQuery) All() ContactList {
	if cq.conn = NewDBConn(); cq.conn != nil {
		defer cq.conn.Close()
	}

	err := cq.fillUsers()
	if err != nil {
		log.Print(err)
		return NewContactList(0)
	}
	return cq.collection
}

func (cq *ContactQuery) fillUsers() (err error) {
	ps, err := cq.selectUsersStmt()
	if err != nil {
		return
	}
	defer ps.Close()

	rows, err := ps.Query(cq.limit, cq.offset)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		contact := NewContact()
		rows.Scan(
			&contact.Id,
			&contact.Email,
			&contact.FirstName,
			&contact.LastName,
			&contact.MiddleName,
			&contact.DateOfBirth,
			&contact.Sex,
		)

		cq.collection = append(cq.collection, contact)
	}

	return
}

func (cq *ContactQuery) selectUsersStmt() (*sql.Stmt, error) {
	if cq.conn == nil {
		return nil, errors.New("Can't connect to DB")
	}

	return cq.conn.Prepare(`
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
      limit $1
      offset $2`)
}
