package main

import (
	"bytes"
	"log"
	"strconv"
)

func NewContactList(size int) *ContactList {
	return &ContactList{i: -1, list: make([]*Contact, 0, size)}
}

type ContactList struct {
	i    int
	list []*Contact
}

func (c *ContactList) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteRune('[')
	listSize := len(c.list)
	for i := 0; i < listSize; i++ {
		data, _ := c.list[i].MarshalJSON()
		buf.Write(data)
		if i < listSize-1 {
			buf.WriteRune(',')
		}
	}
	buf.WriteRune(']')

	return buf.Bytes(), nil
}

func (c *ContactList) Next() *Contact {
	c.i++

	if len(c.list) < c.i+1 {
		c.i--
		return nil
	}

	return c.list[c.i]
}

func (c *ContactList) Append(contact *Contact) {
	c.list = append(c.list, contact)
}

func (c ContactList) Ids() []byte {
	ids := bytes.NewBuffer([]byte("{"))

	strings := make([][]byte, 0)
	for _, next := range c.list {
		if id, err := next.Id.Value(); err == nil {
			strings = append(strings, []byte(strconv.FormatInt(id.(int64), 10)))
		}
	}
	result := bytes.Join(strings, []byte(","))

	if _, err := ids.Write(result); err != nil {
		log.Print(err)
	}

	if _, err := ids.Write([]byte("}")); err != nil {
		log.Print(err)
		return []byte("{}")
	}

	return ids.Bytes()
}

func (c ContactList) Any() bool {
	return len(c.list) > 0
}

func (c ContactList) Items() []*Contact {
	return c.list
}
