package services

import "github.com/AlexanderSchiff/rest/models"

// PetServicePrototype is the interface for the pet service
type PetServicePrototype interface {
	Create(pet models.Pet) (models.Pet, error)
	Update(pet models.Pet) (models.Pet, error)
	GetByPetID(id int64) (models.Pet, error)
	GetByStatus(status string) ([]models.Pet, error)
	GetByTags(tags []string) ([]models.Pet, error)
	Delete(id int64) (string, error)
}