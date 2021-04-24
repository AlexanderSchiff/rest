package models

// Address is a struct containing address details.
type Address struct {
	ID int64 `gorm:"primary_key"`
	Street string `json:"street"`
	City string `json:"city"`
	State string `json:"state"`
	Zip string `json:"zip"`
}