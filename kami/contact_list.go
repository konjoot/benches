package main

func NewContactList(size int) ContactList {
	return ContactList{i: -1, Items: make([]*Contact, 0, size)}
}

type ContactList struct {
	i     int
	Items []*Contact
}

func (c *ContactList) Next() *Contact {
	c.i++

	if len(c.Items) < c.i+1 {
		c.i--
		return nil
	}

	return c.Items[c.i]
}

func (c ContactList) Any() bool {
	return len(c.Items) > 0
}

// func (c ContactList) Ids() []int {
// 	arr := make([]int, len(c.Items))

// 	for i, contact := range c.Items {
// 		arr[i] = contact.Id
// 	}

// 	return arr
// }

// func (c ContactList) IntIds() []int {
// 	arr := make([]int, len(c.Items))

// 	for i, contact := range c.Items {
// 		arr[i] = int(contact.Id)
// 	}

// 	return arr
// }

// func (cl *ContactList) Append(c Contact) {
// 	cl.Items = append(cl.Items, &c)
// }
