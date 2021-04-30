package controllers

import (
	"net/http"
	"github.com/AlexanderSchiff/rest/models"
	"github.com/AlexanderSchiff/rest/services"
	"github.com/gin-gonic/gin"
)

// UserController is the /user controller
type UserController struct {
	Router *gin.Engine
	UserService services.UserServicePrototype
}

// Handle is the /user endpoint
func (uC UserController) Handle() {
	uC.Router.POST("/user", uC.CreateUser)
	uC.Router.POST("/user/createWithList", uC.CreateUsersWithListInput)
	uC.Router.GET("/user/:username", uC.GetUserByName)
	uC.Router.PUT("/user/:username", uC.UpdateUser)
	uC.Router.DELETE("/user/:username", uC.DeleteUser)
}

// CreateUser creates a new user
func (uC UserController) CreateUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users := make([]models.User, 1)
	users[1] = user

	users, err := uC.UserService.Create(users)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": users[0]})
}

// CreateUsersWithListInput creates multiple users
func (uC UserController) CreateUsersWithListInput(context *gin.Context) {
	var users []models.User
	if err := context.ShouldBindJSON(&users); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := uC.UserService.Create(users)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	context.JSON(http.StatusCreated, gin.H{"data": users})
}

// GetUserByName get a user by name
func (uC UserController) GetUserByName(context *gin.Context) {
	var user models.User
	username := context.Param("username")
	user, err := uC.UserService.GetByUsername(username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user})
}

// UpdateUser updates an existing user
func (uC UserController) UpdateUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uC.UserService.Update(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser deletes a user
func (uC UserController) DeleteUser(context *gin.Context) {
	username := context.Param("username")
	message, err := uC.UserService.Delete(username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": message})
}