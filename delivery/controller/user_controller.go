package controller

import (
	"ecommerce/config"
	"ecommerce/entity"
	"ecommerce/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUC usecase.UserUseCase
	rg     *gin.RouterGroup
}

func (u *UserController) createHandler(ctx *gin.Context) {
	var payload entity.User
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := u.userUC.RegisterNewUser(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (u *UserController) listHandle(c *gin.Context) {
	users, err := u.userUC.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (u *UserController) updatedHandle(c *gin.Context) {
	dataId := c.Param("id")

	if dataId == "" {
		c.JSON(http.StatusBadRequest, "invalid ID")
		return
	}

	var data entity.User

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	data.ID = dataId

	updatedUser, err := u.userUC.Update(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to update user: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updatedUser})
}
func (u *UserController) deleteHandle(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, "invalid ID")
		return
	}

	_, err := u.userUC.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to update user: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, "Deleted Succesfully")
}

func (u *UserController) Route() {
	u.rg.POST(config.UserPost, u.createHandler)
	u.rg.GET(config.UserList, u.listHandle)
	u.rg.PUT(config.UserUpdate, u.updatedHandle)
	u.rg.DELETE(config.UserDelete, u.deleteHandle)
}

func NewUserController(userUC usecase.UserUseCase, rg *gin.RouterGroup) *UserController {
	return &UserController{
		userUC: userUC,
		rg:     rg,
	}
}
