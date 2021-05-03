package models

// Order in a store
type Order struct {
	ID int64 `json:"id" gorm:"primary_key"`
	PetID int64 `json:"petId"`
	Quantity int `json:"quantity"`
	ShipDate string `json:"shipDate"`
	Status OrderStatus `json:"status"`
	Complete bool `json:"complete"`
}

// OrderStatus is an enum
type OrderStatus string
const (
	placed = "placed"
	approved = "approved"
	delivered = "delivered"
)
