package main

import (
	// "database/sql"
	"gopkg.in/guregu/null.v3"
)

func NewProfile() *Profile {
	return &Profile{}
}

type Profile struct {
	Id        null.Int    `json:",omitempty"`
	Type      null.String `json:",omitempty"`
	Subjects  []*Subject  `json:",omitempty"`
	ClassUnit ClassUnit   `json:",omitempty"`
	School    School      `json:",omitempty"`
}

// func (p *Profile) MarshalJSON() ([]byte, error) {
// 	return MarshalJSON(p)
// }
