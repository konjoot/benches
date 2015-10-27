package main

import (
	"database/sql"
)

func NewSubject() *Subject {
	return &Subject{}
}

type Subject struct {
	Id   sql.NullInt64
	Name sql.NullString
}

func (s *Subject) MarshalJSON() ([]byte, error) {
	return MarshalJSON(s)
}
