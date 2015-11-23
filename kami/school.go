package main

import (
	// "bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

type School struct {
	Id   sql.NullInt64
	Name sql.NullString
	Guid sql.NullString
}

func (s *School) MarshalJSON() ([]byte, error) {
	buf := NewBuffer()
	defer bufferPool.Put(buf)

	buf.WriteString(`{"id":`)
	buf.WriteString(strconv.FormatInt(s.Id.Int64, 10))
	if s.Name.Valid {
		buf.WriteString(`,"name":`)
		data, _ := json.Marshal(&s.Name.String)
		buf.Write(data)
	}
	if s.Guid.Valid {
		buf.WriteString(`,"guid":"`)
		buf.WriteString(s.Guid.String)
		buf.WriteRune('"')
	}
	buf.WriteRune('}')

	return buf.Bytes(), nil
}

func (s School) Value() (driver.Value, error) {
	if s.Id.Valid || s.Name.Valid || s.Guid.Valid {
		return &s, nil
	}

	return nil, nil
}
