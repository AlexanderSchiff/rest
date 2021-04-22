package main

// Order in a store
type Order struct {
	ID int64
	PetID int64
	Quantity int
	ShipDate string
	Status OrderStatus
	Complete bool
}

// OrderStatus is an enum
type OrderStatus string
const (
	placed = "placed"
	approved = "approved"
	delivered = "delivered"
)
