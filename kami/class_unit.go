package main

import (
	"database/sql"
	"database/sql/driver"
	"strconv"
	"time"

	"github.com/lib/pq"
)

type ClassUnit struct {
	Id         sql.NullInt64
	Name       sql.NullString
	EnlistedOn pq.NullTime
	LeftOn     pq.NullTime
}

func (cu *ClassUnit) MarshalJSON() ([]byte, error) {
	buf := NewBuffer()
	defer bufferPool.Put(buf)

	buf.WriteString(`{"id":`)
	buf.WriteString(strconv.FormatInt(cu.Id.Int64, 10))
	if cu.Name.Valid {
		buf.WriteString(`,"name":"`)
		buf.WriteString(cu.Name.String)
		buf.WriteRune('"')
	}
	if cu.EnlistedOn.Valid {
		buf.WriteString(`,"enlisted_on":"`)
		buf.WriteString(cu.EnlistedOn.Time.Format(time.RFC3339))
		buf.WriteRune('"')
	}
	if cu.LeftOn.Valid {
		buf.WriteString(`,"left_on":"`)
		buf.WriteString(cu.LeftOn.Time.Format(time.RFC3339))
		buf.WriteRune('"')
	}
	buf.WriteRune('}')

	return buf.Bytes(), nil
}

func (cu ClassUnit) Value() (driver.Value, error) {
	if cu.Id.Valid || cu.Name.Valid || cu.EnlistedOn.Valid || cu.LeftOn.Valid {
		return &cu, nil
	}

	return nil, nil
}
