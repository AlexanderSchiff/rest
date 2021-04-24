package controllers

import (
  "github.com/gin-gonic/gin"
)

// StoreController /store endpoint
func StoreController(router *gin.Engine) {
	router.GET("/store/inventory", func(c *gin.Context) {
		// tbd
	})
	router.POST("/store/order", func(c *gin.Context) {
		// tbd
	})
	router.GET("/store/order/:orderId", func(c *gin.Context) {
		// tbd
	})
	router.DELETE("/store/order/:orderId", func(c *gin.Context) {
		// tbd
	})
}