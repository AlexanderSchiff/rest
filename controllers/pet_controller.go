package controllers

import (
	"sync"
	"github.com/gin-gonic/gin"
)

var waitGroup sync.WaitGroup

// PetController is the /pet route
func PetController(router *gin.Engine) {
	router.PUT("/pet", UpdatePet)
	router.POST("/pet", AddPet)
	router.GET("/pet/findByStatus", FindPetsByStatus)
	router.GET("/pet/findByTags", FindPetsByTags)
	router.GET("/pet/:petId", GetPetByID)
	router.POST("/pet/:petId", UpdatePetWithForm)
	router.DELETE("/pet/:petId", DeletePet)
	router.POST("/pet/:petId/uploadImage", UploadFile)
}

// UpdatePet update an existing pet
func UpdatePet(context *gin.Context) {

}

// AddPet add a new pet to the store
func AddPet(context *gin.Context) {

}

// FindPetsByStatus finds pets by status
func FindPetsByStatus(context *gin.Context) {

}

// FindPetsByTags finds pets by tags
func FindPetsByTags(context *gin.Context) {

}

// GetPetByID gets a pet by ID
func GetPetByID(context *gin.Context) {

}

// UpdatePetWithForm updates a pet using a form
func UpdatePetWithForm(context *gin.Context) {

}

// DeletePet deletes a pet
func DeletePet(context *gin.Context) {

}

// UploadFile uploads a picture of a pet
func UploadFile(context *gin.Context) {

}