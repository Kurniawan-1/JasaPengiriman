package usecase

import (
	"encoding/json"
	"errors"
	"jasaPengiriman/models"

	"github.com/gin-gonic/gin"
)

type BarangUseCase struct {
	barangRepo models.BarangRepo
}

func NewBarangUseCase(barangRepo models.BarangRepo) *BarangUseCase {
	return &BarangUseCase{
		barangRepo,
	}
}

func (barangUC *BarangUseCase) GetAllBarang() []models.Barang {
	result := barangUC.barangRepo.GetAllBarang()
	if result == nil {
		return nil
	}
	return result
}

func (barangUC *BarangUseCase) CreateBarang(c *gin.Context) error {
	var result models.Barang
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = barangUC.barangRepo.CreateBarang(&result)
	if err != nil {
		return err
	}

	return nil
}

func (barangUC *BarangUseCase) UpdateBarang(c *gin.Context) error {
	var result models.Barang
	id := c.Param("id")
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = barangUC.barangRepo.UpdateBarang(&result, id)
	if err != nil {
		return err
	}
	return nil
}

func (barangUC *BarangUseCase) DeleteBarang(c *gin.Context) error {
	var input struct {
		Id json.Number
	}
	c.ShouldBindJSON(&input)
	id, _ := input.Id.Int64()

	if id <= 0 {
		return errors.New("ID Tidak ada")
	}

	err := barangUC.barangRepo.DeleteBarang(id)
	if err != nil {
		return err
	}

	return nil
}
