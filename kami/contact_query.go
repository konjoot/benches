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
	collection ContactList
	conn       *sql.DB
}

func (cq *ContactQuery) All() []*Contact {
	if cq.conn = NewDBConn(); cq.conn != nil {
		defer cq.conn.Close()
	}

	if ok := cq.fillUsers(); !ok {
		return NewContactList(0).Items()
	}

	if err := cq.fillDependentData(); err != nil {
		log.Print(err)
	}

	return cq.collection.Items()
}

func (cq *ContactQuery) fillUsers() (ok bool) {
	var err error

	defer func() {
		if err != nil {
			log.Print(err)
		}

		if cq.collection.Any() {
			ok = true
		}
	}()

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

		cq.collection.Append(contact)
	}

	return
}

func (cq *ContactQuery) fillDependentData() (err error) {
	ps, err := cq.selectDependentDataStmt()
	if err != nil {
		return
	}
	defer ps.Close()

	rows, err := ps.Query(cq.collection.Ids())
	if err != nil {
		return
	}
	defer rows.Close()

	var userId sql.NullInt64
	var current *Contact

	if ok := cq.collection.Next(); ok {
		current = cq.collection.Current()
	}

	for rows.Next() {

		profile := NewProfile()
		rows.Scan(
			&profile.Id,
			&profile.Type,
			&userId,
		)

		for current.Id != userId {

			if ok := cq.collection.Next(); !ok {
				break
			}
			current = cq.collection.Current()
		}

		current.Profiles = append(current.Profiles, profile)
	}

	return
}

func (cq *ContactQuery) selectUsersStmt() (*sql.Stmt, error) {
	if cq.conn == nil {
		return nil, errors.New("Can't connect to DB")
	}

	return cq.conn.Prepare(`
		select	id,
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

func (cq *ContactQuery) selectDependentDataStmt() (*sql.Stmt, error) {
	if cq.conn == nil {
		return nil, errors.New("Can't connect to DB")
	}

	return cq.conn.Prepare(`
		select	id,
						type,
						user_id
			from profiles
			where deleted_at is null
				and user_id = any($1::integer[])
			order by user_id, id`)
}
