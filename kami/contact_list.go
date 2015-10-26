package main

import (
	"bytes"
	"log"
	"strconv"
)

func NewContactList(size int) ContactList {
	return ContactList{i: -1, list: make([]*Contact, 0, size)}
}

type ContactList struct {
	i    int
	list []*Contact
}

func (c *ContactList) Next() bool {
	c.i++

	if len(c.list) < c.i+1 {
		c.i--
		return false
	}

	return true
}

func (c ContactList) Current() *Contact {
	if c.i < 0 {
		return nil
	}

	if len(c.list) < c.i+1 {
		return nil
	}

	return c.list[c.i]
}

func (c *ContactList) Append(contact *Contact) {
	c.list = append(c.list, contact)
}

func (c ContactList) Ids() string {
	ids := bytes.NewBuffer([]byte(`{`))

	strings := make([][]byte, 0)
	for c.Next() {
		if id, err := c.Current().Id.Value(); err == nil {
			strings = append(strings, []byte(strconv.FormatInt(id.(int64), 10)))
		}
	}
	result := bytes.Join(strings, []byte(`,`))

	if _, err := ids.Write(result); err != nil {
		log.Print(err)
	}

	if _, err := ids.Write([]byte(`}`)); err != nil {
		log.Print(err)
		return "{}"
	}

	return ids.String()
}

func (c ContactList) Any() bool {
	return len(c.list) > 0
}

func (c ContactList) Items() []*Contact {
	return c.list
}

func (c *ContactList) First() *Contact {
	c.Next()
	return c.Current()
}
