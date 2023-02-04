package handler

import (
	"jasaPengiriman/models"
	"jasaPengiriman/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type pelangganHandler struct {
	pelangganUC models.PelangganUseCase
}

func PelangganRoute(pelangganUC models.PelangganUseCase, r *gin.RouterGroup) {
	uc := pelangganHandler{
		pelangganUC,
	}
	v2 := r.Group("pelanggan")
	v2.GET("", uc.GetAllPelanggan)
	v2.POST("", uc.CreatePelanggan)
	v2.PUT(":id", uc.UpdatePelanggan)
	v2.DELETE("", uc.DeletePelanggan)
}

func (pelangganHandler *pelangganHandler) GetAllPelanggan(c *gin.Context) {
	result := pelangganHandler.pelangganUC.GetAllPelanggan()
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

func (pelangganHandler *pelangganHandler) CreatePelanggan(c *gin.Context) {
	err := pelangganHandler.pelangganUC.CreatePelanggan(c)
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

func (pelangganHandler *pelangganHandler) UpdatePelanggan(c *gin.Context) {
	err := pelangganHandler.pelangganUC.UpdatePelanggan(c)
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

func (pelangganHandler *pelangganHandler) DeletePelanggan(c *gin.Context) {
	err := pelangganHandler.pelangganUC.DeletePelanggan(c)
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
