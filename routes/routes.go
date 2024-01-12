package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitRouter(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	e.GET("/products", func(c echo.Context) error {
		return c.String(http.StatusOK, "Products")
	})

	e.POST("/products", func(c echo.Context) error {
		return c.String(http.StatusOK, "Product created")
	})

	e.GET("/votes", func(c echo.Context) error {
		return c.String(http.StatusOK, "Votes")
	})

	e.POST("/votes", func(c echo.Context) error {
		return c.String(http.StatusOK, "Vote created")
	})
}
