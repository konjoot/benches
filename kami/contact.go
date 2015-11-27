package main

import "time"

func NewContact() *Contact {
	return &Contact{}
}

type Contact struct {
	Id          *int32     `json:"id,omitempty"`
	Email       *string    `json:"email,omitempty"`
	FirstName   *string    `json:"firstName,omitempty"`
	LastName    *string    `json:"lastName,omitempty"`
	MiddleName  *string    `json:"middleName,omitempty"`
	DateOfBirth *time.Time `json:"dateOfBirth,omitempty"`
	Sex         *int32     `json:"sex,omitempty"`
	Profiles    []*Profile `json:"profiles,omitempty"`
}

func (c *Contact) LastProfile() *Profile {
	count := len(c.Profiles)

	if count > 0 {
		return c.Profiles[count-1]
	}

	return nil
}
