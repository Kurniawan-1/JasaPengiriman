package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"jasaPengiriman/models"

	"github.com/gin-gonic/gin"
)

type PelangganUseCase struct {
	pelangganRepo models.PelangganRepo
}

func NewPelangganUseCase(pelangganRepo models.PelangganRepo) *PelangganUseCase {
	return &PelangganUseCase{
		pelangganRepo,
	}
}

func (pelangganUC *PelangganUseCase) GetAllPelanggan() []models.Pelanggan {
	result := pelangganUC.pelangganRepo.GetAllPelanggan()
	if result == nil {
		return nil
	}
	return result
}

func (pelangganUC *PelangganUseCase) CreatePelanggan(c *gin.Context) error {
	var result models.Pelanggan
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = pelangganUC.pelangganRepo.CreatePelanggan(&result)
	if err != nil {
		return err
	}

	return nil
}

func (pelangganUC *PelangganUseCase) UpdatePelanggan(c *gin.Context) error {
	var result models.Pelanggan
	id := c.Param("id")
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = pelangganUC.pelangganRepo.UpdatePelanggan(&result, id)
	if err != nil {
		return err
	}
	return nil
}

func (pelangganUC *PelangganUseCase) DeletePelanggan(c *gin.Context) error {
	var input struct {
		Id json.Number
	}

	c.ShouldBindJSON(&input)

	id, _ := input.Id.Int64()
	fmt.Println(id)

	if id <= 0 {
		return errors.New("ID Tidak ada")
	}

	err := pelangganUC.pelangganRepo.DeletePelanggan(id)
	if err != nil {
		return err
	}

	return nil
}
