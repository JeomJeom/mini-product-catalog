package controller

import (
	"net/http"
	"thalesapi/data/dto"
	"thalesapi/data/params"
	"thalesapi/db/builder"
	"thalesapi/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	svc service.ProductService
}

func NewProductController(svc service.ProductService) *ProductController {
	return &ProductController{svc: svc}
}

func (p ProductController) CreateProduct(c *gin.Context) {
	var product dto.ProductMutableAttrs
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := p.svc.AddProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save product"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (p ProductController) GetAllProducts(c *gin.Context) {
	productFilter := params.NewProductFilter(c)
	paginated, err := builder.NewPaginatedRequest(c, "updated_at")
	products, err := p.svc.GetProducts(productFilter, paginated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (p ProductController) GetProduct(c *gin.Context) {
	id := c.Param("id")

	product, err := p.svc.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch product"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (p ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var modifiedProduct dto.ProductMutableAttrs
	if err := c.BindJSON(&modifiedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	products, err := p.svc.UpdateProduct(id, &modifiedProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (p ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := p.svc.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.Status(http.StatusNoContent) // 204
}
