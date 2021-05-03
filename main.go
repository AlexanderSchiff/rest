package main

import (
	"sync"
	"github.com/AlexanderSchiff/rest/controllers"
	"github.com/AlexanderSchiff/rest/models"
	"github.com/AlexanderSchiff/rest/services"
	"github.com/gin-gonic/gin"
)

// Main function
func main() {
	models.ConnectDataBase()
	dB := models.DB
	var waitGroup sync.WaitGroup
	waitGroup.Add(3)
	router := gin.Default()
	pet := func() {
		petService := services.PetService{DB: dB}
		petController := controllers.PetController{Router: router, PetService: petService}
		petController.Handle()
		waitGroup.Done()
	}
	store := func() {
		storeService := services.StoreService{DB: dB}
		storeController := controllers.StoreController{Router: router, StoreService: storeService}
		storeController.Handle()
		waitGroup.Done()
	}
	user := func() {
		userService := services.UserService{DB: dB}
		userController := controllers.UserController{Router: router, UserService: userService}
		userController.Handle()
		waitGroup.Done()
	}
	go pet()
	go store()
	go user()
	waitGroup.Wait()
	router.Run()
}