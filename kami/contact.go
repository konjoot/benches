package main

import (
	"bytes"
	// "fmt"
	"sync"
	"time"
)

// const QUOTE = "\""

const QUOTE = byte(34)

var jsPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

// var bufPool = sync.Pool{
// 	New: func() interface{} {
// 		return new(bytes.Buffer)
// 	},
// }

func NewContact() *Contact {
	return &Contact{
		Id:         NewInteger(),
		Email:      NewString(),
		FirstName:  NewString(),
		LastName:   NewString(),
		MiddleName: NewString(),
		Sex:        NewInteger(),
	}
}

type Contact struct {
	Id          Integer    `json:"id,omitempty" db:"id"`
	Email       String     `json:"email,omitempty" db:"email"`
	FirstName   String     `json:"firstName,omitempty" db:"first_name"`
	LastName    String     `json:"lastName,omitempty" db:"last_name"`
	MiddleName  String     `json:"middleName,omitempty" db:"middle_name"`
	DateOfBirth *time.Time `json:"dateOfBirth,omitempty" db:"date_of_birth"`
	Sex         Integer    `json:"sex,omitempty" db:"sex"`
	Profiles    []*Profile `json:"profiles,omitempty" db:"-"`
}

func (c *Contact) LastProfile() *Profile {
	count := len(c.Profiles)

	if count > 0 {
		return c.Profiles[count-1]
	}

	return nil
}

func NewString() String {
	return String(make([]byte, 0, 257))
}

type String []byte

func (s *String) Set(val []byte) {
	*s = (*s)[0 : len(val)+2]
	copy((*s)[1:len(val)+1], val)
}

func (s String) MarshalJSON() ([]byte, error) {
	s[0] = QUOTE
	s[len(s)-1] = QUOTE
	return s, nil
}

func NewInteger() Integer {
	return Integer(make([]byte, 0, 255))
}

type Integer []byte

func (i *Integer) Set(val []byte) {
	*i = (*i)[0:len(val)]
	copy(*i, val)
}

func (i Integer) MarshalJSON() ([]byte, error) {
	return i, nil
}

// func (c *Contact) SetId(val []byte) {
// 	buf := bufPool.Get().(*bytes.Buffer)
// 	buf.Reset()
// 	buf.Write(val)
// 	if c.Id == nil {
// 		c.Id
// 	}
// 	c.Id = buf.Bytes()
// 	bufPool.Put(buf)
// }

// func (c *Contact) SetEmail(val []byte) {
// 	buf := bufPool.Get().(*bytes.Buffer)
// 	buf.Reset()
// 	quote := []byte("\"")
// 	buf.Write(quote)
// 	buf.Write(val)
// 	buf.Write(quote)
// 	c.Email = buf.Bytes()
// 	bufPool.Put(buf)
// }

// func (c *Contact) SetFirstName(val []byte) {
// 	buf := bufPool.Get().(*bytes.Buffer)
// 	buf.Reset()
// 	quote := []byte("\"")
// 	buf.Write(quote)
// 	buf.Write(val)
// 	buf.Write(quote)
// 	c.FirstName = buf.Bytes()
// 	bufPool.Put(buf)
// }

// func (c *Contact) SetLastName(val []byte) {
// 	buf := bufPool.Get().(*bytes.Buffer)
// 	buf.Reset()
// 	quote := []byte("\"")
// 	buf.Write(quote)
// 	buf.Write(val)
// 	buf.Write(quote)
// 	c.LastName = buf.Bytes()
// 	bufPool.Put(buf)
// }

// func (c *Contact) SetMiddleName(val []byte) {
// 	buf := bufPool.Get().(*bytes.Buffer)
// 	buf.Reset()
// 	quote := []byte("\"")
// 	buf.Write(quote)
// 	buf.Write(val)
// 	buf.Write(quote)
// 	c.MiddleName = buf.Bytes()
// 	bufPool.Put(buf)
// }

// func (c *Contact) SetSex(val []byte) {
// 	buf := bufPool.Get().(*bytes.Buffer)
// 	buf.Reset()
// 	buf.Write(val)
// 	c.Sex = buf.Bytes()
// 	bufPool.Put(buf)
// }

// func (_ *Contact) SetString(str *String, val []byte) {
// 	if str == nil {
// 		str = new(String)
// 	}
// 	str.Bytes.Reset()
// 	quote := []byte("\"")
// 	str.Bytes.Write(quote)
// 	str.Bytes.Write(val)
// 	str.Bytes.Write(quote)
// }

// func (_ *Contact) SetInteger(i *Integer, val []byte) {
// 	if i == nil {
// 		i = new(Integer)
// 	}
// 	i.Bytes.Reset()
// 	i.Bytes.Write(val)
// }
