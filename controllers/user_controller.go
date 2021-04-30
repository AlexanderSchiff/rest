package controllers

import (
	"net/http"

	"github.com/AlexanderSchiff/rest/models"
	"github.com/AlexanderSchiff/rest/services"
	"github.com/gin-gonic/gin"
)

// UserController is the /user endpoint
func UserController(router *gin.Engine, userService services.UserServicePrototype) {
	router.POST("/user", userService.CreateUser)
	router.POST("/user/createWithList", userService.CreateUsersWithListInput)
	router.GET("/user/:username", userService.GetUserByName)
	router.PUT("/user/:username", userService.UpdateUser)
	router.DELETE("/user/:username", userService.DeleteUser)
}

// CreateUser creates a new user
func CreateUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&user)
	context.JSON(http.StatusCreated, gin.H{"data": user})
}

// CreateUsersWithListInput creates multiple users
func CreateUsersWithListInput(context *gin.Context) {
	var users []models.User
	if err:= context.ShouldBindJSON(&users); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&users)
	context.JSON(http.StatusCreated, gin.H{"data": users})
}

// GetUserByName get a user by name
func GetUserByName(context *gin.Context) {
	var user models.User

	if err := models.DB.Where("username = ?", context.Param("username")).First(&user).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
	}

	context.JSON(http.StatusOK, gin.H{"data": user})
}

// UpdateUser updates an existing user
func UpdateUser(context *gin.Context) {

}

// DeleteUser deletes a user
func DeleteUser(context *gin.Context) {

}