package models

// Pet is the pet struct
type Pet struct {
	ID int64 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Category Category `json:"category"`
	PhotoUrls []string `json:"photoUrls"`
	Tags []Tag `json:"tags"`
	Status string `json:"status"`
}