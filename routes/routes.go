package routes

import (
	"backend/controller"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {

	productHandler := controller.ProductHandler{Database: db}

	e.GET(
		"/",
		func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, Echo!")
		},
	)

	e.GET(
		"/products",
		productHandler.GetProducts,
	)

	e.POST(
		"/products",
		productHandler.CreateProduct,
	)

	e.GET(
		"/votes",
		func(c echo.Context) error {
			return c.String(http.StatusOK, "Votes")
		},
	)

	e.POST(
		"/votes",
		func(c echo.Context) error {
			return c.String(http.StatusOK, "Vote created")
		},
	)
}
