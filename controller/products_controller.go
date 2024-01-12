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

// GetProducts @Summary Get all products
// @Description Get all products
// @ID get-all-products
// @Produce  json
// @Success 200 {array} models.Product
// @Router /products [get]
func (h *ProductHandler) GetProducts(c echo.Context) error {
	var products []models.Product
	result := h.Database.Find(&products)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Something went wrong")
	}
	return c.JSON(http.StatusOK, products)
}

// CreateProduct @Summary Create a new product
// @Description Create a new product with the input payload
// @ID create-product
// @Accept  json
// @Produce  json
// @Param product body models.Product true "Create product"
// @Success 201 {object} models.Product
// @Router /products [post]
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
