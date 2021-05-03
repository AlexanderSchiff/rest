package services

import "github.com/AlexanderSchiff/rest/models"

// UserServicePrototype is the interface that contains methods for the UserService
type UserServicePrototype interface {
	Create([]models.User) ([]models.User, error)
	GetByUsername(username string) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(username string) (string, error)
}