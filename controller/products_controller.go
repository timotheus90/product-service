package controller

import (
	"backend/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type ProductHandler struct {
	Database *gorm.DB
}

func (h *ProductHandler) GetProducts(c echo.Context) error {
	var products []models.Product
	result := h.Database.Find(&products)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}
	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	product := models.Product{}
	if err := c.Bind(&product); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request payload")
	}
	result := h.Database.Create(&product)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}
	return c.JSON(http.StatusCreated, product)
}
