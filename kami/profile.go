package main

import (
	"bytes"
	"database/sql"
	"strconv"
)

func NewProfile() *Profile {
	return &Profile{}
}

type Profile struct {
	Id        sql.NullInt64
	Type      sql.NullString
	Subjects  []*Subject
	ClassUnit ClassUnit
	School    School
}

// func (p *Profile) MarshalJSON() ([]byte, error) {
// 	return MarshalJSON(p)
// }

func (p *Profile) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(`{"id":`)
	buf.WriteString(strconv.FormatInt(p.Id.Int64, 10))
	if p.Type.Valid {
		buf.WriteString(`,"type":"`)
		buf.WriteString(p.Type.String)
		buf.WriteRune('"')
	}
	buf.WriteString(`,"subjects":[`)
	subjectsCount := len(p.Subjects)
	for i := 0; i < subjectsCount; i++ {
		data, _ := p.Subjects[i].MarshalJSON()
		buf.Write(data)
		if i < subjectsCount-1 {
			buf.WriteRune(',')
		}
	}
	buf.WriteRune(']')
	if p.Type.String == "StudentProfile" {
		buf.WriteString(`,"class_unit":`)
		data, _ := p.ClassUnit.MarshalJSON()
		buf.Write(data)
	}
	buf.WriteString(`,"school":`)
	data, _ := p.School.MarshalJSON()
	buf.Write(data)
	buf.WriteRune('}')

	return buf.Bytes(), nil
}
