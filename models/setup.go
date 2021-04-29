package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

// DB database
var DB *gorm.DB

// ConnectDataBase connect database
func ConnectDataBase() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if (err != nil) {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Address{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Customer{})
	DB = db
}