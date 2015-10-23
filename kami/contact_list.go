package main

func NewContactList() ContactList {
	return make(ContactList, 0, 100)
}

type ContactList []Contact
