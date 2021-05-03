package controllers

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/AlexanderSchiff/rest/services"
	"github.com/AlexanderSchiff/rest/models"
)

// PetController is the /store controller
type PetController struct {
	Router *gin.Engine
	PetService services.PetServicePrototype
}

// Handle is the /pet route
func (pC PetController) Handle() {
	pC.Router.PUT("/pet", pC.UpdatePet)
	pC.Router.POST("/pet", pC.AddPet)
	pC.Router.GET("/pet/findByStatus", pC.FindPetsByStatus)
	pC.Router.GET("/pet/findByTags", pC.FindPetsByTags)
	pC.Router.GET("/pet/:petId", pC.GetPetByID)
	pC.Router.PUT("/pet/:petId", pC.UpdatePetWithForm)
	pC.Router.DELETE("/pet/:petId", pC.DeletePet)
}

// UpdatePet update an existing pet
func (pC PetController) UpdatePet(context *gin.Context) {
	var pet models.Pet
	if err := context.ShouldBindJSON(&pet); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pet, err := pC.PetService.Update(pet)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": pet})
}

// AddPet add a new pet to the store
func (pC PetController) AddPet(context *gin.Context) {
	var pet models.Pet
	if err := context.ShouldBindJSON(&pet); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pet, err := pC.PetService.Create(pet)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": pet})
}

// FindPetsByStatus finds pets by status
func (pC PetController) FindPetsByStatus(context *gin.Context) {
	status := context.Query("status")
	if status != "available" && status != "pending" && status != "sold" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
		return
	}

	pets, err := pC.PetService.GetByStatus(status)
		if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": pets})
}

// FindPetsByTags finds pets by tags
func (pC PetController) FindPetsByTags(context *gin.Context) {
	// find by tags. need to bind to multiple query params
}

// GetPetByID gets a pet by ID
func (pC PetController) GetPetByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("petId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pet, err := pC.PetService.GetByPetID(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": pet})
}

// UpdatePetWithForm updates a pet using a form
func (pC PetController) UpdatePetWithForm(context *gin.Context) {

}

// DeletePet deletes a pet
func (pC PetController) DeletePet(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("petID"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	message, err := pC.PetService.Delete(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": message})
}