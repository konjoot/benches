package main

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"

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
}

func (c Contact) MarshalJSON() (res []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			if ee, ok := e.(error); ok {
				err = ee
			}
		}
	}()

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

	jsonData := bytes.Join(result, []byte(","))
	if _, err = buffer.Write(jsonData); err != nil {
		return
	}

	if _, err = buffer.Write([]byte(`}`)); err != nil {
		return
	}

	return buffer.Bytes(), err
}
