package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"backend/models"
)

func main() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.Product{}, &models.Vote{})
	if err != nil {
		panic("failed to migrate database")
	}

	db.Create(&models.Product{ID: 1, Name: "Product 1", Description: "Description 1"})

	var product models.Product
	db.First(&product, 1)
	fmt.Println(product)

	db.Model(&product).Update("Name", "Product 1 updated")

	db.Delete(&product)
}
