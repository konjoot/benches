package main

import (
	// "database/sql"
	"gopkg.in/guregu/null.v3"
)

func NewSubject() *Subject {
	return &Subject{}
}

type Subject struct {
	Id   null.Int    `json:",omitempty"`
	Name null.String `json:",omitempty"`
}

// func (s *Subject) MarshalJSON() ([]byte, error) {
// 	return MarshalJSON(s)
// }
