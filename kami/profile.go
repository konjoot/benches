package main

import (
	"database/sql"
)

func NewProfile() *Profile {
	return &Profile{}
}

type Profile struct {
	Id       sql.NullInt64
	Type     sql.NullString
	Subjects []*Subject
	ClassUnit
	School
}

func (p *Profile) MarshalJSON() ([]byte, error) {
	return MarshalJSON(p)
}
