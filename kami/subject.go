package main

func NewSubject() *Subject {
	return &Subject{}
}

type Subject struct {
	Id   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
