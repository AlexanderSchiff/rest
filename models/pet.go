package models

// Pet is the pet struct
type Pet struct {
	ID int64 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Category Category `json:"category" gorm:"foreignKey:PetID"`
	PhotoUrls []string `json:"photoUrls"`
	Tags []Tag `json:"tags" gorm:"foreignKey:PetID"`
	Status string `json:"status"`
	Orders []Order `gorm:"foreignKey:PetID"`
}