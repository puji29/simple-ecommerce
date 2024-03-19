package controller

import (
	"ecommerce/config"
	"ecommerce/entity"
	"ecommerce/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUC usecase.ProductUseCase
	rg        *gin.RouterGroup
}

func (p *ProductController) createHandler(c *gin.Context) {
	var payload entity.Product

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	product, err := p.productUC.CreateNewProduct(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func (p *ProductController) listProductHandler(c *gin.Context) {
	product, err := p.productUC.FindAllProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func (p *ProductController) updateProductHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	var product entity.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	product.ID = id

	updateProduct, err := p.productUC.UpdateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updateProduct})
}
func (p *ProductController) deleteProductHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	_, err := p.productUC.DeletedProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Delete succesfully")
}
func (p *ProductController) getProductByIdHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	product, err := p.productUC.FindProductById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}
func (p *ProductController) getProductByProductNameHandler(c *gin.Context) {
	productName := c.Param("productName")
	if productName == "" {
		c.JSON(http.StatusBadRequest, "invalid product name")
		return
	}
	product, err := p.productUC.FindProductByProductName(productName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func (p *ProductController) Route() {
	p.rg.POST(config.ProductPost, p.createHandler)
	p.rg.GET(config.ProductGet, p.listProductHandler)
	p.rg.PUT(config.ProductPut, p.updateProductHandler)
	p.rg.DELETE(config.ProductDelete, p.deleteProductHandler)
	p.rg.GET(config.ProductGetById, p.getProductByIdHandler)
	p.rg.GET(config.ProductGetByProductName, p.getProductByProductNameHandler)
}

func NewProductController(productUC usecase.ProductUseCase, rg *gin.RouterGroup) *ProductController {
	return &ProductController{productUC: productUC, rg: rg}
}
