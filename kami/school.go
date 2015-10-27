package main

import (
	"database/sql"
	"database/sql/driver"
)

type School struct {
	Id   sql.NullInt64
	Name sql.NullString
	Guid sql.NullString
}

func (s *School) MarshalJSON() ([]byte, error) {
	return MarshalJSON(s)
}

func (s School) Value() (driver.Value, error) {
	if s.Id.Valid || s.Name.Valid || s.Guid.Valid {
		return &s, nil
	}

	return nil, nil
}
