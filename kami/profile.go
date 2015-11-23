package main

import ()

func NewProfile() *Profile {
	return &Profile{}
}

type Profile struct {
	Id        *int32
	Type      *string
	Subjects  []*Subject
	ClassUnit ClassUnit
	School    School
}

// func (p *Profile) MarshalJSON() ([]byte, error) {
// 	return MarshalJSON(p)
// }
