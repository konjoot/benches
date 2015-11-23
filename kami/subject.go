package main

import ()

func NewSubject() *Subject {
	return &Subject{}
}

type Subject struct {
	Id   *int32
	Name *string
}

// func (s *Subject) MarshalJSON() ([]byte, error) {
// 	return MarshalJSON(s)
// }
