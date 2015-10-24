package main

func NewContactList(size int) ContactList {
	return make(ContactList, 0, size)
}

type ContactList []*Contact
