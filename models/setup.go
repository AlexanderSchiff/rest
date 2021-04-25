package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

// DB database
var DB *gorm.DB

// ConnectDataBase connect database
func ConnectDataBase() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if (err != nil) {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Address{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Customer{})
	DB = db
}