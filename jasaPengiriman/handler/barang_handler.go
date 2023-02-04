package handler

import (
	"jasaPengiriman/models"
	"jasaPengiriman/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type barangHandler struct {
	barangUC models.BarangUseCase
}

func BarangRoute(barangUC models.BarangUseCase, r *gin.RouterGroup) {
	uc := barangHandler{
		barangUC,
	}
	v2 := r.Group("barang")
	v2.GET("", uc.GetAllBarang)
	v2.POST("", uc.CreateBarang)
	v2.PUT(":id", uc.UpdateBarang)
	v2.DELETE("", uc.DeleteBarang)
}

func (barangHandler *barangHandler) GetAllBarang(c *gin.Context) {
	result := barangHandler.barangUC.GetAllBarang()
	if result == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Data:    result,
		Status:  "success",
		Message: "success",
		Meta:    nil,
	})
}

func (barangHandler *barangHandler) CreateBarang(c *gin.Context) {
	err := barangHandler.barangUC.CreateBarang(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		Status:  "success",
		Message: "success",
	})
}

func (barangHandler *barangHandler) UpdateBarang(c *gin.Context) {
	err := barangHandler.barangUC.UpdateBarang(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "success",
		Message: "success",
	})
}

func (barangHandler *barangHandler) DeleteBarang(c *gin.Context) {
	err := barangHandler.barangUC.DeleteBarang(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "success",
		Message: "success",
	})
}
