package services

import (
	"github.com/AlexanderSchiff/rest/models"
	"gorm.io/gorm"
)

// UserService is the user service
type UserService struct {
	DB *gorm.DB
}

// Create calls DB.create
func (uS UserService) Create(users []models.User) (outputUsers []models.User, err error) {
	uS.DB.Create(&users)
	outputUsers = users
	return outputUsers, nil
}

// GetByUsername retrieves a user by username
func (uS UserService) GetByUsername(username string) (user models.User, err error) {
	err = uS.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

// Update updates an existing user
func (uS UserService) Update(user models.User) (output models.User, err error) {
	err = uS.DB.Where("username = ?", user.Username).First(&output).Error
	if err != nil {
		return output, err
	}

	uS.DB.Save(user)
	output = user
	return output, err
}

// Delete delete an existing user
func (uS UserService) Delete(username string) (message string, err error) {
	var user models.User
	err = uS.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return "not found!", err
	}

	uS.DB.Delete(&user)
	return "success", err
}