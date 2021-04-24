package models

// Customer represents a customer on an main.Order
type Customer struct {
	ID int64 `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Address Address `json:"address"`
}