package handler

import (
	"jasaPengiriman/models"
	"jasaPengiriman/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type pengirimanHandler struct {
	pengirimanUC models.PengirimanUseCase
}

func PengirimanRoute(pengirimanUC models.PengirimanUseCase, r *gin.RouterGroup) {
	uc := pengirimanHandler{
		pengirimanUC,
	}
	v2 := r.Group("pengiriman")
	v2.GET("", uc.GetAllPengiriman)
	v2.POST("", uc.CreatePengiriman)
	v2.PUT(":id", uc.UpdatePengiriman)
	v2.DELETE("", uc.DeletePengiriman)
}

func (pengirimanHandler *pengirimanHandler) GetAllPengiriman(c *gin.Context) {
	result := pengirimanHandler.pengirimanUC.GetAllPengiriman()
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

func (pengirimanHandler *pengirimanHandler) CreatePengiriman(c *gin.Context) {
	err := pengirimanHandler.pengirimanUC.CreatePengiriman(c)
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

func (pengirimanHandler *pengirimanHandler) UpdatePengiriman(c *gin.Context) {
	err := pengirimanHandler.pengirimanUC.UpdatePengiriman(c)
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

func (pengirimanHandler *pengirimanHandler) DeletePengiriman(c *gin.Context) {
	err := pengirimanHandler.pengirimanUC.DeletePengiriman(c)
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
