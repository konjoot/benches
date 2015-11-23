package main

import (
	"database/sql"
	"strconv"

	"github.com/lib/pq"
)

func NewContact() *Contact {
	return &Contact{}
}

type Contact struct {
	Id          sql.NullInt64
	Email       sql.NullString
	FirstName   sql.NullString
	LastName    sql.NullString
	MiddleName  sql.NullString
	DateOfBirth pq.NullTime
	Sex         sql.NullInt64
	Profiles    []*Profile
}

func (c *Contact) MarshalJSON() ([]byte, error) {
	buf := NewBuffer()
	defer bufferPool.Put(buf)

	buf.WriteString(`{"id":`)
	buf.WriteString(strconv.FormatInt(c.Id.Int64, 10))
	if c.Email.Valid {
		buf.WriteString(`,"email":"`)
		buf.WriteString(c.Email.String)
		buf.WriteRune('"')
	}
	if c.FirstName.Valid {
		buf.WriteString(`,"first_name":"`)
		buf.WriteString(c.FirstName.String)
		buf.WriteRune('"')
	}
	if c.LastName.Valid {
		buf.WriteString(`,"last_name":"`)
		buf.WriteString(c.LastName.String)
		buf.WriteRune('"')
	}
	if c.MiddleName.Valid {
		buf.WriteString(`,"middle_name":"`)
		buf.WriteString(c.MiddleName.String)
		buf.WriteRune('"')
	}
	if c.DateOfBirth.Valid {
		buf.WriteString(`,"date_of_birth":"`)
		buf.WriteString(c.DateOfBirth.Time.Format("2006-01-02"))
		buf.WriteRune('"')
	}
	buf.WriteString(`,"profiles":[`)
	profilesCount := len(c.Profiles)
	for i := 0; i < profilesCount; i++ {
		data, _ := c.Profiles[i].MarshalJSON()
		buf.Write(data)
		if i < profilesCount-1 {
			buf.WriteRune(',')
		}
	}
	buf.WriteRune(']')
	buf.WriteRune('}')

	return buf.Bytes(), nil
}

func (c *Contact) LastProfile() *Profile {
	count := len(c.Profiles)

	if count > 0 {
		return c.Profiles[count-1]
	}

	return nil
}
