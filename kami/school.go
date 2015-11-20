package main

import (
	// "database/sql"
	// "database/sql/driver"
	"gopkg.in/guregu/null.v3"
)

type School struct {
	Id   null.Int    `json:",omitempty"`
	Name null.String `json:",omitempty"`
	Guid null.String `json:",omitempty"`
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
