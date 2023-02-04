package usecase

import (
	"encoding/json"
	"errors"
	"jasaPengiriman/models"

	"github.com/gin-gonic/gin"
)

type KurirUseCase struct {
	kurirRepo models.KurirRepo
}

func NewKurirUseCase(kurirRepo models.KurirRepo) *KurirUseCase {
	return &KurirUseCase{
		kurirRepo,
	}
}

func (kurirUC *KurirUseCase) GetAllKurir() []models.Kurir {
	result := kurirUC.kurirRepo.GetAllKurir()
	if result == nil {
		return nil
	}
	return result
}

func (kurirUC *KurirUseCase) CreateKurir(c *gin.Context) error {
	var result models.Kurir
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = kurirUC.kurirRepo.CreateKurir(&result)
	if err != nil {
		return err
	}

	return nil
}

func (kurirUC *KurirUseCase) UpdateKurir(c *gin.Context) error {
	var result models.Kurir
	id := c.Param("id")
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = kurirUC.kurirRepo.UpdateKurir(&result, id)
	if err != nil {
		return err
	}
	return nil
}

func (kurirUC *KurirUseCase) DeleteKurir(c *gin.Context) error {
	var input struct {
		Id json.Number
	}
	c.ShouldBindJSON(&input)
	id, _ := input.Id.Int64()

	if id <= 0 {
		return errors.New("ID Tidak ada")
	}

	err := kurirUC.kurirRepo.DeleteKurir(id)
	if err != nil {
		return err
	}

	return nil
}
