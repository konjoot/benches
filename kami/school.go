package main

import ()

func NewSchool() *School {
	return &School{}
}

type School struct {
	Id   *int32
	Name *string
	Guid *string
}

// func (s *School) MarshalJSON() ([]byte, error) {
// 	return MarshalJSON(s)
// }

// func (s School) Value() (driver.Value, error) {
// 	if s.Id.Valid || s.Name.Valid || s.Guid.Valid {
// 		return &s, nil
// 	}

// 	return nil, nil
// }
