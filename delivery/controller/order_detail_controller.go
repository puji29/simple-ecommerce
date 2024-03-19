package controller

import (
	"ecommerce/config"
	"ecommerce/entity"
	"ecommerce/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderDetailController struct {
	odUc usecase.OrderDetailUsecase
	rg   *gin.RouterGroup
}

func (o *OrderDetailController) createHandle(c *gin.Context) {
	var orderDetail entity.OrderDetail

	if err := c.ShouldBind(&orderDetail); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newOrderDetail, err := o.odUc.InsertNew(orderDetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": newOrderDetail})
}

func (o *OrderDetailController) Route() {
	o.rg.POST(config.OrderDetailInsert, o.createHandle)
}

func NewOrderDetailController(odUc usecase.OrderDetailUsecase, rg *gin.RouterGroup) *OrderDetailController {
	return &OrderDetailController{odUc, rg}
}
