package main

import (
// "bytes"
// "log"
// "strconv"
)

func NewContactList(size int) ContactList {
	return ContactList{i: -1, list: make([]*Contact, 0, size)}
}

type ContactList struct {
	i    int
	list []*Contact
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
	// c.Ids = append(c.Ids, *contact.Id)
	c.list = append(c.list, contact)
}

// func (c ContactList) Ids() []byte {
// 	ids := bytes.NewBuffer([]byte("{"))

// 	strings := make([][]byte, 0)
// 	for _, next := range c.list {
// 		id := *next.Id
// 		strings = append(strings, []byte(strconv.FormatInt(int64(id), 10)))
// 	}
// 	result := bytes.Join(strings, []byte(","))

// 	if _, err := ids.Write(result); err != nil {
// 		log.Print(err)
// 	}

// 	if _, err := ids.Write([]byte("}")); err != nil {
// 		log.Print(err)
// 		return []byte("{}")
// 	}

// 	return ids.Bytes()
// }

func (c *ContactList) Ids() []int32 {
	arr := make([]int32, len(c.list))

	for i, contact := range c.list {
		arr[i] = *contact.Id
	}

	return arr
}

func (c ContactList) Any() bool {
	return len(c.list) > 0
}

func (c ContactList) Items() []*Contact {
	return c.list
}
