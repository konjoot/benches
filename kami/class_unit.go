package main

import (
	"time"
)

func NewClassUnit() *ClassUnit {
	return &ClassUnit{}
}

type ClassUnit struct {
	Id         *int32
	Name       *string
	EnlistedOn *time.Time `json:,omitempty`
	LeftOn     *time.Time `json:,omitempty`
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
