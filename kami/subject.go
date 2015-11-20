package main

import (
	"bytes"
	"database/sql"
	"strconv"
)

func NewSubject() *Subject {
	return &Subject{}
}

type Subject struct {
	Id   sql.NullInt64
	Name sql.NullString
}

// func (s *Subject) MarshalJSON() ([]byte, error) {
// 	return MarshalJSON(s)
// }

func (s *Subject) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(`{"id":`)
	buf.WriteString(strconv.FormatInt(s.Id.Int64, 10))
	if s.Name.Valid {
		buf.WriteString(`,"name":"`)
		buf.WriteString(s.Name.String)
		buf.WriteRune('"')
	}
	buf.WriteRune('}')

	return buf.Bytes(), nil
}
