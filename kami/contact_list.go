package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	// "strconv"
	"sync"
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
	contact.SetIndex(len(c.list) - 1)
	c.list = append(c.list, contact)
}

// func (c ContactList) Ids() []byte {
// 	ids := bytes.NewBuffer([]byte("{"))

// 	strings := make([][]byte, 0)
// 	for _, next := range c.list {
// 		strings = append(strings, []byte(strconv.FormatInt(next.GetId(), 10)))
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

func (c ContactList) Ids() (res []int32) {
	for _, next := range c.list {
		res = append(res, *next.Id)
	}

	return
}

func (c ContactList) Any() bool {
	return len(c.list) > 0
}

func (c ContactList) Items() []*Contact {
	return c.list
}

func (c ContactList) JSON(w io.Writer) {
	var wg sync.WaitGroup
	out := make(chan []byte, len(c.list))
	defer wg.Wait()
	defer close(out)

	for _, contact := range c.list {
		wg.Add(1)
		go func(c *Contact) {
			defer wg.Done()

			res, err := json.Marshal(c)

			if err != nil {
				log.Print(err)
			}

			out <- res

		}(contact)
	}

	openSquare := []byte("[")
	closeSquare := []byte("]")
	comma := []byte(",")
	buf := make([][]byte, 0)
	w.Write(openSquare)

	for i := 0; i < len(c.list); i++ {
		select {
		case v := <-out:
			buf = append(buf, v)
		}
	}

	w.Write(bytes.Join(buf, comma))
	w.Write(closeSquare)

	// return result, nil
}
