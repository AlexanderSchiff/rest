package models

// Pet is the pet struct
type Pet struct {
	ID int64 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Category Category `json:"category" gorm:"foreignKey:PetID"`
	Photos []Photo `gorm:"foreignKey:PetID"`
	Tags []Tag `json:"tags" gorm:"foreignKey:PetID"`
	Status PetStatus `json:"status"`
	Orders []Order `gorm:"foreignKey:PetID"`
}

// PetStatus is an enum
type PetStatus string
const (
	available = "available"
	pending = "pending"
	sold = "sold"
)