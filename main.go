package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

// Main function
func main() {
	router := gin.Default()

	// /pet
	router.PUT("/pet", func(c *gin.Context) {
		// tbd
	})
	router.POST("/pet", func(c *gin.Context) {
		// tbd
	})
	router.GET("/pet/findByStatus", func(c *gin.Context) {
		// tbd
	})
	router.GET("/pet/findByTags", func(c *gin.Context) {
		// tbd
	})
	router.GET("/pet/:petId", func(c *gin.Context) {
		// tbd
	})
	router.POST("/pet/:petId", func(c *gin.Context) {
		// tbd
	})
	router.DELETE("/pet/:petId", func(c *gin.Context) {
		// tbd
	})
	router.POST("/pet/:petId/uploadImage", func(c *gin.Context) {
		// tbd
	})

	// /store
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

	// /user
	router.POST("/user", func(c *gin.Context) {
		// tbd
	})
	router.POST("/user/createWithList", func(c *gin.Context) {
		// tbd
	})
	router.GET("/user/login", func(c *gin.Context) {
		// tbd
	})
	router.GET("/user/logout", func(c *gin.Context) {
		// tbd
	})
	router.GET("/user/:username", func(c *gin.Context) {
		// tbd
	})
	router.PUT("/user/:username", func(c *gin.Context) {
		// tbd
	})
	router.DELETE("/user/:username", func(c *gin.Context) {
		// tbd
	})
}