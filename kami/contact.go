package main

import (
	"time"
)

func NewContact() *Contact {
	return &Contact{}
}

type Contact struct {
	Id          *int32
	Email       *string
	FirstName   *string
	LastName    *string
	MiddleName  *string
	DateOfBirth *time.Time
	Sex         *int32
	Profiles    []*Profile
}

// func (c *Contact) MarshalJSON() ([]byte, error) {
// 	return MarshalJSON(c)
// }

func (c *Contact) LastProfile() *Profile {
	count := len(c.Profiles)

	if count > 0 {
		return c.Profiles[count-1]
	}

	return nil
}
