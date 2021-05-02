package services

import (
	"github.com/AlexanderSchiff/rest/models"
	"gorm.io/gorm"
)

// StoreService is the user service
type StoreService struct {
	DB *gorm.DB
}

// Create calls DB.create
func (sS StoreService) Create(order models.Order) (outputOrder models.Order, err error) {
	sS.DB.Create(&order)
	outputOrder = order
	return outputOrder, nil
}

// GetByOrderID retrieves an order by id
func (sS StoreService) GetByOrderID(id int64) (outputOrder models.Order, err error) {
	err = sS.DB.Where("id = ?", id).First(&outputOrder).Error
	return outputOrder, err
}

// Delete delete an existing order
func (sS StoreService) Delete(id int64) (message string, err error) {
	var order models.Order
	err = sS.DB.Where("id = ?", id).First(&order).Error
	if err != nil {
		return "not found!", err
	}

	sS.DB.Delete(&order)
	return "success", err
}