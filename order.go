package main

// Order in a store
type Order struct {
	id int64
	petID int64
	quantity int
	shipDate string
	status OrderStatus
	complete bool
}

// OrderStatus is an enum
type OrderStatus string
const (
	placed = "placed"
	approved = "approved"
	delivered = "delivered"
)
