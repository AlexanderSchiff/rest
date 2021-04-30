package services

import "github.com/gin-gonic/gin"

// UserServicePrototype is the interface that contains methods for the UserService
type UserServicePrototype interface {
	CreateUser(context *gin.Context)
	CreateUsersWithListInput(context *gin.Context)
	GetUserByName(context *gin.Context)
	UpdateUser(context *gin.Context)
	DeleteUser(context *gin.Context)
}