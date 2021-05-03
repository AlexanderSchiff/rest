package models

// User represents a customer
type User struct {
	ID int64 `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	UserStatus int `json:"userStatus"`
}