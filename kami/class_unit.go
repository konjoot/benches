package main

import (
	// "database/sql"
	// "database/sql/driver"

	// "github.com/lib/pq"
	"gopkg.in/guregu/null.v3"
)

type ClassUnit struct {
	Id         null.Int    `json:",omitempty"`
	Name       null.String `json:",omitempty"`
	EnlistedOn null.Time   `json:",omitempty"`
	LeftOn     null.Time   `json:",omitempty"`
}

// func (cu *ClassUnit) MarshalJSON() ([]byte, error) {
// 	return MarshalJSON(cu)
// }

// func (cu ClassUnit) Value() (driver.Value, error) {
// 	if cu.Id.Valid || cu.Name.Valid || cu.EnlistedOn.Valid || cu.LeftOn.Valid {
// 		return &cu, nil
// 	}

// 	return nil, nil
// }
