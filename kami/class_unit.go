package main

import (
	"time"
)

type ClassUnit struct {
	Id         *int32
	Name       *string
	EnlistedOn *time.Time
	LeftOn     *time.Time
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
