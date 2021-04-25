package controllers

import (
  "github.com/gin-gonic/gin"
)

// StoreController /store endpoint
func StoreController(router *gin.Engine) {
	router.GET("/store/inventory", GetInventory)
	router.POST("/store/order", PlaceOrder)
	router.GET("/store/order/:orderId", GetOrderByID)
	router.DELETE("/store/order/:orderId", DeleteOrder)
}

// GetInventory gets the inventory of a store
func GetInventory(context *gin.Context) {

}

// PlaceOrder places a new order in a store
func PlaceOrder(context *gin.Context) {

}

// GetOrderByID gets the order by ID
func GetOrderByID(context *gin.Context) {
}

// DeleteOrder deletes an order
func DeleteOrder(context *gin.Context) {

}