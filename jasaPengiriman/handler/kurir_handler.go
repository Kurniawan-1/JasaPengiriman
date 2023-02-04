package handler

import (
	"jasaPengiriman/models"
	"jasaPengiriman/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type kurirHandler struct {
	kurirUC models.KurirUseCase
}

func KurirRoute(kurirUC models.KurirUseCase, r *gin.RouterGroup) {
	uc := kurirHandler{
		kurirUC,
	}
	v2 := r.Group("kurir")
	v2.GET("", uc.GetAllKurir)
	v2.POST("", uc.CreateKurir)
	v2.PUT(":id", uc.UpdateKurir)
	v2.DELETE("", uc.DeleteKurir)
}

func (kurirHandler *kurirHandler) GetAllKurir(c *gin.Context) {
	result := kurirHandler.kurirUC.GetAllKurir()
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

func (kurirHandler *kurirHandler) CreateKurir(c *gin.Context) {
	err := kurirHandler.kurirUC.CreateKurir(c)
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

func (kurirHandler *kurirHandler) UpdateKurir(c *gin.Context) {
	err := kurirHandler.kurirUC.UpdateKurir(c)
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

func (kurirHandler *kurirHandler) DeleteKurir(c *gin.Context) {
	err := kurirHandler.kurirUC.DeleteKurir(c)
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
