package main

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sync"
)

var bufferPool sync.Pool

type Buffer struct {
	bytes.Buffer
}

func NewBuffer() *Buffer {
	if v := bufferPool.Get(); v != nil {
		b := v.(*Buffer)
		b.Reset()
		return b
	}
	return new(Buffer)
}

func MarshalJSON(i interface{}) (res []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("MarshalJSON error: %s\n", e))
		}
	}()

	buffer := bytes.NewBuffer([]byte("{"))

	rtype := reflect.TypeOf(i).Elem()
	rval := reflect.ValueOf(i).Elem()
	count := rtype.NumField()

	result := make([][]byte, 0)

	for i := range make([]struct{}, count) {
		key := rtype.Field(i)
		value := rval.Field(i)
		if val, ok := value.Interface().(driver.Valuer); ok {
			if v, err := val.Value(); v != nil && err == nil {
				if vJson, err := json.Marshal(v); err == nil {
					res := []byte(`"`)
					res = append(res, []byte(key.Name)...)
					res = append(res, []byte(`":`)...)
					res = append(res, vJson...)
					result = append(result, res)
				}
			}
		} else if value.Kind() == reflect.Slice && value.Len() > 0 {
			if vJson, err := json.Marshal(value.Interface()); err == nil {
				res := []byte(`"`)
				res = append(res, []byte(key.Name)...)
				res = append(res, []byte(`":`)...)
				res = append(res, vJson...)
				result = append(result, res)
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
