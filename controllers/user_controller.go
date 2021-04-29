package controllers

import (
  "github.com/gin-gonic/gin"
)

// UserController is the /user endpoint
func UserController(router *gin.Engine) {
	router.POST("/user", CreateUser)
	router.POST("/user/createWithList", CreateUsersWithListInput)
	router.GET("/user/:username", GetUserByName)
	router.PUT("/user/:username", UpdateUser)
	router.DELETE("/user/:username", DeleteUser)
}

// CreateUser creates a new user
func CreateUser(context *gin.Context) {

}

// CreateUsersWithListInput creates multiple users
func CreateUsersWithListInput(context *gin.Context) {

}

// GetUserByName get a user by name
func GetUserByName(context *gin.Context) {

}

// UpdateUser updates an existing user
func UpdateUser(context *gin.Context) {

}

// DeleteUser deletes a user
func DeleteUser(context *gin.Context) {

}