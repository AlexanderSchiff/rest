package controllers

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/AlexanderSchiff/rest/models"
	"github.com/AlexanderSchiff/rest/services"
)

// StoreController is the /store controller
type StoreController struct {
	Router *gin.Engine
	StoreService services.StoreServicePrototype
}

// Handle /store endpoint
func (sC StoreController) Handle() {
	sC.Router.POST("/store/order", sC.PlaceOrder)
	sC.Router.GET("/store/order/:orderId", sC.GetOrderByID)
	sC.Router.DELETE("/store/order/:orderId", sC.DeleteOrder)
}

// PlaceOrder places a new order in a store
func (sC StoreController) PlaceOrder(context *gin.Context) {
	var order models.Order
	if err := context.ShouldBindJSON(&order); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := sC.StoreService.Create(order)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": order})
}

// GetOrderByID gets the order by ID
func (sC StoreController) GetOrderByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("orderId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := sC.StoreService.GetByOrderID(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": order})
}

// DeleteOrder deletes an order
func (sC StoreController) DeleteOrder(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("orderId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	message, err := sC.StoreService.Delete(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": message})
}