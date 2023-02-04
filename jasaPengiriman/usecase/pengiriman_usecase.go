package usecase

import (
	"encoding/json"
	"errors"
	"jasaPengiriman/models"
	"time"

	"github.com/gin-gonic/gin"
)

type PengirimanUseCase struct {
	pengirimanRepo models.PengirimanRepo
}

func NewPengirimanUseCase(pengirimanRepo models.PengirimanRepo) *PengirimanUseCase {
	return &PengirimanUseCase{
		pengirimanRepo,
	}
}

func (pengirimanUC *PengirimanUseCase) GetAllPengiriman() []models.Pengiriman {
	result := pengirimanUC.pengirimanRepo.GetAllPengiriman()
	if result == nil {
		return nil
	}
	return result
}

func (pengirimanUC *PengirimanUseCase) CreatePengiriman(c *gin.Context) error {
	var result models.Pengiriman
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	layout := "02/01/2006 MST"
	waktuKirim, _ := time.Parse(layout, result.TanggalPengiriman)
	result.TanggalKirim = waktuKirim

	if result.TanggalKirim.Before(time.Now()) {
		return errors.New("tanggal kirim tidak valid")
	}

	waktuTerima, _ := time.Parse(layout, result.TanggalPenerimaan)
	result.TanggalTerima = waktuTerima

	if result.TanggalTerima.Before(result.TanggalKirim) {
		return errors.New("tanggal terima tidak valid")
	}

	err = pengirimanUC.pengirimanRepo.CreatePengiriman(&result)
	if err != nil {
		return err
	}

	return nil
}

func (pengirimanUC *PengirimanUseCase) UpdatePengiriman(c *gin.Context) error {
	var result models.Pengiriman
	id := c.Param("id")
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = pengirimanUC.pengirimanRepo.UpdatePengiriman(&result, id)
	if err != nil {
		return err
	}
	return nil
}

func (pengirimanUC *PengirimanUseCase) DeletePengiriman(c *gin.Context) error {
	var input struct {
		Id json.Number
	}
	c.ShouldBindJSON(&input)
	id, _ := input.Id.Int64()

	if id <= 0 {
		return errors.New("ID Tidak ada")
	}

	err := pengirimanUC.pengirimanRepo.DeletePengiriman(id)
	if err != nil {
		return err
	}

	return nil
}
