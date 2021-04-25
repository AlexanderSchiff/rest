package main

import (
	"sync"

	"github.com/AlexanderSchiff/rest/controllers"
	"github.com/AlexanderSchiff/rest/models"
	"github.com/gin-gonic/gin"
)

// Main function
func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(4)
	router := gin.Default()
	pet := func() {
		controllers.PetController(router)
		waitGroup.Done()
	}
	store := func() {
		controllers.StoreController(router)
		waitGroup.Done()		
	}
	user := func() {
		controllers.UserController(router)
		waitGroup.Done()
	}
	db := func() {
		models.ConnectDataBase()
		waitGroup.Done()
	}
	go pet()
	go store()
	go user()
	go db()
	waitGroup.Wait()
}