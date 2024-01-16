package controller

import (
	"ecommerce/config"
	"ecommerce/entity"
	"ecommerce/usecase"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ImageController struct {
	imageUC usecase.ImageUseCase
	rg      *gin.RouterGroup
}

func (i *ImageController) createHandler(c *gin.Context) {
	var payload entity.Images

	image, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NO image"})
		return
	}
	fileName := fmt.Sprintf("%d%s", time.Now().Unix(), filepath.Ext(image.Filename))
	uploadDir := "./uploads"
	filepath := filepath.Join(uploadDir, fileName)

	if err := c.SaveUploadedFile(image, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return
	}
	payload.ID = uuid.New().String()
	payload.Image = fileName

	cretedImage, err := i.imageUC.CreateImage(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save image"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cretedImage, "message": "file upload succesfully"})

}
func (i *ImageController) listHandler(c *gin.Context) {
	image, err := i.imageUC.FindAllImage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": image})

}
func (i *ImageController) updatedHandler(c *gin.Context) {

	id := c.Param("id")
	existingImage, err := i.imageUC.ImageById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "image id not found")
		return
	}

	image, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, "no image")
		return
	}

	fileName := fmt.Sprintf("%d%s", time.Now().Unix(), filepath.Ext(image.Filename))

	uploadDir := "./uploads"
	filepath := filepath.Join(uploadDir, fileName)

	if err := c.SaveUploadedFile(image, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, "failed to save file")
		return
	}
	existingImage.ID = uuid.New().String()
	existingImage.Image = fileName

	updateImage, err := i.imageUC.ImageUpdated(existingImage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updateImage, "message": "updated succcesfully"})
}

func (i *ImageController) GetByIdHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	image, err := i.imageUC.ImageById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": image})
}
func (i *ImageController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	_, err := i.imageUC.DeleteI(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "delete Succesfully"})
}

func (i *ImageController) Route() {
	i.rg.POST(config.ImageInsert, i.createHandler)
	i.rg.GET(config.ImageGet, i.listHandler)
	i.rg.PUT(config.ImagePut, i.updatedHandler)
	i.rg.GET(config.ImageById, i.GetByIdHandler)
	i.rg.DELETE(config.ImageDelete, i.deleteHandler)
}

func NewImageController(imageUC usecase.ImageUseCase, rg *gin.RouterGroup) *ImageController {
	return &ImageController{imageUC, rg}
}
