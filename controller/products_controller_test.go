package controller

import (
	"backend/models"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ProductHandlerTestSuite struct {
	suite.Suite
	Echo        *echo.Echo
	DB          *gorm.DB
	ProductCtrl *ProductHandler
}

func (suite *ProductHandlerTestSuite) SetupTest() {
	suite.Echo = echo.New()
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	err := db.AutoMigrate(&models.Product{})
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.DB = db
	suite.ProductCtrl = &ProductHandler{Database: db}
}

func (suite *ProductHandlerTestSuite) TestGetProducts() {
	mockProduct := models.Product{Name: "Mock Product", Description: "Mock Description"}
	suite.DB.Create(&mockProduct)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := suite.Echo.NewContext(req, rec)

	if assert.NoError(suite.T(), suite.ProductCtrl.GetProducts(c)) {
		assert.Equal(suite.T(), http.StatusOK, rec.Code)
		var products []models.Product
		err := json.Unmarshal(rec.Body.Bytes(), &products)
		if err != nil {
			suite.T().Fatal(err)
		}
		assert.Equal(suite.T(), mockProduct.Name, products[0].Name)
		assert.Equal(suite.T(), mockProduct.Description, products[0].Description)
	}
}

func (suite *ProductHandlerTestSuite) TestCreateProduct() {
	mockProduct := models.Product{Name: "Mock Product", Description: "Mock Description"}
	mockProductJSON, _ := json.Marshal(mockProduct)

	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(mockProductJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.Echo.NewContext(req, rec)

	if assert.NoError(suite.T(), suite.ProductCtrl.CreateProduct(c)) {
		assert.Equal(suite.T(), http.StatusCreated, rec.Code)
		var product models.Product
		err := json.Unmarshal(rec.Body.Bytes(), &product)
		if err != nil {
			suite.T().Fatal(err)
		}
		assert.Equal(suite.T(), mockProduct.Name, product.Name)
		assert.Equal(suite.T(), mockProduct.Description, product.Description)
	}
}

func TestProductHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(ProductHandlerTestSuite))
}
