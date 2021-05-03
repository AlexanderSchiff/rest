package models

// Photo contains photo info
type Photo struct {
	ID int64 `gorm:"primary_key"`
	URL string
	PetID int64
}