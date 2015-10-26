package main

import (
	"database/sql"

	"github.com/lib/pq"
)

func NewContact() *Contact {
	return &Contact{}
}

type Contact struct {
	Id          sql.NullInt64
	Email       sql.NullString
	FirstName   sql.NullString
	LastName    sql.NullString
	MiddleName  sql.NullString
	DateOfBirth pq.NullTime
	Sex         sql.NullInt64
	Profiles    []*Profile
}

func (c *Contact) MarshalJSON() ([]byte, error) {
	return MarshalJSON(c)
}
