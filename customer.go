package main

import (
	"time"
)

type Customer struct {
	Title          string     `json:",omitempty"`
	Initials       string     `json:",omitempty"`
	Name           string     `json:",omitempty"`
	Firstname      string     `json:",omitempty"`
	Surname        string     `json:",omitempty"`
	DOB            *time.Time `json:",omitempty"`
	Address        *Address   `json:",omitempty"`
	Phone          string     `json:",omitempty"`
	Fax            string     `json:",omitempty"`
	Mobile         string     `json:",omitempty"`
	Email          string     `json:",omitempty"`
	InvoiceAddress *Address   `json:",omitempty"`
}

type Address struct {
	Address  []string `json:",omitempty"`
	Postcode string   `json:",omitempty"`
}
