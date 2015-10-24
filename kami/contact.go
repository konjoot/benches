package main

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"github.com/lib/pq"
	"reflect"
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
}

// todo: add panic recover
func (c Contact) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte("{"))

	rtype := reflect.TypeOf(&c).Elem()
	rval := reflect.ValueOf(&c).Elem()
	count := rtype.NumField()

	result := make([][]byte, 0)

	for i := range make([]struct{}, count) {
		key := rtype.Field(i)
		value := rval.Field(i).Interface()
		if val, ok := value.(driver.Valuer); ok {
			if v, err := val.Value(); v != nil && err == nil {
				if vJson, err := json.Marshal(v); err == nil {
					res := []byte(`"`)
					res = append(res, []byte(key.Name)...)
					res = append(res, []byte(`":`)...)
					res = append(res, vJson...)
					result = append(result, res)
				}
			}
		}
	}

	_, err := buffer.Write(bytes.Join(result, []byte(",")))
	if err != nil {
		return nil, err
	}

	if _, err := buffer.Write([]byte(`}`)); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
