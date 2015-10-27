package main

import (
	"database/sql"
	"database/sql/driver"

	"github.com/lib/pq"
)

type ClassUnit struct {
	Id         sql.NullInt64
	Name       sql.NullString
	EnlistedOn pq.NullTime
	LeftOn     pq.NullTime
}

func (cu *ClassUnit) MarshalJSON() ([]byte, error) {
	return MarshalJSON(cu)
}

func (cu ClassUnit) Value() (driver.Value, error) {
	if cu.Id.Valid || cu.Name.Valid || cu.EnlistedOn.Valid || cu.LeftOn.Valid {
		return &cu, nil
	}

	return nil, nil
}
