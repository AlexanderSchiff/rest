package models

// Tag for pet
type Tag struct {
	ID int64 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	PetID int64
}