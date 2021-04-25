package models

// Category of pet
type Category struct {
	ID int64 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}