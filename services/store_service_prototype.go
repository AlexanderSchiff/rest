package services

import "github.com/AlexanderSchiff/rest/models"

// StoreServicePrototype is the interface for the store service
type StoreServicePrototype interface {
	Create(store models.Order) (models.Order, error)
	GetByOrderID(id int64) (models.Order, error)
	Delete(id int64) (string, error)
}