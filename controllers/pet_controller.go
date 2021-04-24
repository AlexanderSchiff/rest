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
	router.GET("/pet/:petId", GetPetById)
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

func FindPetsByStatus(context *gin.Context) {

}

func FindPetsByTags(context *gin.Context) {

}

func GetPetById(context *gin.Context) {

}

func UpdatePetWithForm(context *gin.Context) {

}

func DeletePet(context *gin.Context) {

}

func UploadFile(context *gin.Context) {

}