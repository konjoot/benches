package main

func NewProfile() *Profile {
	return &Profile{}
}

type Profile struct {
	Id        *int32     `json:"id,omitempty"`
	Type      *string    `json:"type,omitempty"`
	UserId    *int32     `json:"-"`
	ClassUnit *ClassUnit `json:"classUnit,omitempty"`
	School    *School    `json:"school,omitempty"`
	Subjects  []*Subject `json:"subjects,omitempty"`
}
