package main

import ()

func NewProfile() *Profile {
	return &Profile{}
}

type Profile struct {
	Id        *int32
	Type      *string
	Subjects  []*Subject `json:,omitempty`
	ClassUnit *ClassUnit `json:,omitempty`
	School    *School    `json:,omitempty`
}

// func (p *Profile) MarshalJSON() ([]byte, error) {
// 	return MarshalJSON(p)
// }
