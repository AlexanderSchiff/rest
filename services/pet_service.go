package services

import (
	"github.com/AlexanderSchiff/rest/models"
	"gorm.io/gorm"
	"github.com/ahmetb/go-linq/v3"
)

// PetService is the pet service
type PetService struct {
	DB *gorm.DB
}

// Create calls DB.create
func (pS PetService) Create(pet models.Pet) (outputPet models.Pet, err error) {
	pS.DB.Create(&pet)
	outputPet = pet
	return outputPet, nil
}

// Update updates an existing user
func (pS PetService) Update(pet models.Pet) (output models.Pet, err error) {
	err = pS.DB.First(&output, pet.ID).Error
	if err != nil {
		return output, err
	}

	pS.DB.Save(pet)
	output = pet
	return output, err
}

// GetByPetID retrieves a pet by id
func (pS PetService) GetByPetID(id int64) (outputPet models.Pet, err error) {
	err = pS.DB.First(&outputPet, id).Error
	return outputPet, err
}

// GetByStatus retrieves a pet by status
func (pS PetService) GetByStatus(status string) (outputPets []models.Pet, err error) {
	err = pS.DB.Where("status = ?", status).Find(&outputPets).Error
	return outputPets, err
}

// GetByTags retrieves a pet by status
func (pS PetService) GetByTags(tags []string) (outputPets []models.Pet, err error) {
	var outputTags []models.Tag
	// get names from input tags
	err = pS.DB.Where("name IN ?", tags).Find(&outputTags).Error
	if (err != nil) {
		return nil, err
	}

	var ids []int64
	// get pets from ids
	linq.From(outputTags).Select(func(t interface{}) interface{} {
		return t.(models.Tag).PetID
	}).ToSlice(&ids)
	err = pS.DB.Find(&outputPets, ids).Error
	return outputPets, err
}

// Delete delete an existing order
func (pS PetService) Delete(id int64) (message string, err error) {
	var pet models.Pet
	err = pS.DB.First(&pet, id).Error
	if err != nil {
		return "not found!", err
	}

	pS.DB.Delete(&pet)
	return "success", err
}