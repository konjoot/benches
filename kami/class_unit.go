package main

import "time"

func NewClassUnit() *ClassUnit {
	return &ClassUnit{}
}

type ClassUnit struct {
	Id         *int32     `json:"id,omitempty"`
	Name       *string    `json:"name,omitempty"`
	EnlistedOn *time.Time `json:"enlistedOn,omitempty"`
	LeftOn     *time.Time `json:"leftOn,omitempty"`
}
