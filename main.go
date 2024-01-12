package main

import (
	_ "backend/docs"
	"backend/models"
	"backend/routes"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	db.Create(&models.Product{Name: "Product 1", Description: "Description 1"})
	db.Create(&models.Product{Name: "Product 2", Description: "Description 2"})
	db.Create(&models.Product{Name: "Product 3", Description: "Description 3"})
	db.Create(&models.Product{Name: "Product 4", Description: "Description 4"})

	var product models.Product
	db.First(&product, 1)
	fmt.Println(product)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))
	e.GET(
		"/swagger/*",
		echoSwagger.WrapHandler,
	)
	routes.InitRouter(e, db)
	// Start the server
	err = e.Start(":8080")
	if err != nil {
		panic(err)
	}

}
