package main

// Address is a struct containing address details.
type Address struct {
	Street string `json:"street"`
	City string `json:"city"`
	State string `json:"state"`
	Zip string `json:"zip"`
}