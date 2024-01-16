package controller

import (
	"ecommerce/config"
	"ecommerce/entity"
	"ecommerce/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderTableController struct {
	orderTableUC usecase.OrderTableUseCase
	rg           *gin.RouterGroup
}

func (o *OrderTableController) createHandler(c *gin.Context) {
	var payload entity.OrderTable

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	orderTable, err := o.orderTableUC.CreateOrderTable(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orderTable})
}
func (o *OrderTableController) listHandler(c *gin.Context) {
	orderTable, err := o.orderTableUC.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orderTable})
}
func (o *OrderTableController) updateHandle(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	var orderTable entity.OrderTable

	if err := c.ShouldBindJSON(&orderTable); err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	orderTable.ID = id

	updateOrderTable, err := o.orderTableUC.UpdateById(orderTable)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updateOrderTable})
}
func (o *OrderTableController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "invalid Id")
		return
	}
	_, err := o.orderTableUC.DeleteOT(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Delete Succesfully")
}
func (o *OrderTableController) GetById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	orderTable, err := o.orderTableUC.FindByid(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orderTable})
}

func (o *OrderTableController) Route() {
	o.rg.POST(config.OrderTablePost, o.createHandler)
	o.rg.GET(config.OrderTableList, o.listHandler)
	o.rg.PUT(config.OrderTablePut, o.updateHandle)
	o.rg.DELETE(config.OrderTableDelete, o.deleteHandler)
	o.rg.GET(config.OrderTableByID, o.GetById)
}

func NewOrderTableController(orderTableUC usecase.OrderTableUseCase, rg *gin.RouterGroup) *OrderTableController {
	return &OrderTableController{orderTableUC, rg}
}
