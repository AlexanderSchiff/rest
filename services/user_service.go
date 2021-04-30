package services

import (
	"github.com/AlexanderSchiff/rest/models"
	"gorm.io/gorm"
)

// UserService is the user service
type UserService struct {
	DB *gorm.DB
}

func (uS UserService) Create(users []models.User) (outputUsers []models.User, err error) {

}

func (uS UserService) GetByUsername(username string) (user models.User, err error) {

}

func (uS UserService) Update(user models.User) (output models.User, err error) {

}

func (uS UserService) Delete(username string) (message string, err error) {

}