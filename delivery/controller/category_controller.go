package controller

import (
	"ecommerce/config"
	"ecommerce/entity"
	"ecommerce/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryUc usecase.CategoryUseCase
	rg         *gin.RouterGroup
}

func (a *CategoryController) createHandler(c *gin.Context) {
	var category entity.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newCate, err := a.categoryUc.InsertNew(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": newCate})
}

func (a *CategoryController) Route() {
	a.rg.POST(config.CategoryPost, a.createHandler)

}
func NewCategoryController(categoryUc usecase.CategoryUseCase, rg *gin.RouterGroup) *CategoryController {
	return &CategoryController{categoryUc, rg}
}
