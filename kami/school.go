package main

func NewSchool() *School {
	return &School{}
}

type School struct {
	Id   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Guid *string `json:"guid,omitempty"`
}
