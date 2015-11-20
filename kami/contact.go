package main

import (
	// "database/sql"
	"time"

	// "github.com/lib/pq"
	// "gopkg.in/guregu/null.v3"
)

func NewContact() *Contact {
	return &Contact{}
}

type Contact struct {
	Id          *int32
	Email       *string
	FirstName   *string
	LastName    *string
	MiddleName  *string    `json:",omitempty"`
	DateOfBirth *time.Time `json:",omitempty"`
	Sex         *int32     `json:",omitempty"`
	Profiles    []*Profile `json:",omitempty"`

	index int
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

func (c *Contact) SetIndex(i int) {
	c.index = i
}

func (c *Contact) GetId() int32 {
	return *c.Id
}
