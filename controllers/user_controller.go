package controllers

import (
  "github.com/gin-gonic/gin"
)

// UserController is the /user endpoint
func UserController(router *gin.Engine) {
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