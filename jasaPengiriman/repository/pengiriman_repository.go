package repository

import (
	"jasaPengiriman/models"

	"gorm.io/gorm"
)

type PengirimanRepo struct {
	db *gorm.DB
}

func NewPengirimanRepo(db *gorm.DB) *PengirimanRepo {
	return &PengirimanRepo{db}
}

func (pengirimanRepo *PengirimanRepo) GetAllPengiriman() []models.Pengiriman {
	var result []models.Pengiriman
	data := pengirimanRepo.db.Preload("Pelanggan").Preload("Kurir").Preload("Barang").Find(&result)

	if data.Error != nil {
		return nil
	}

	return result
}

func (pengirimanRepo *PengirimanRepo) CreatePengiriman(data *models.Pengiriman) error {
	err := pengirimanRepo.db.Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (pengirimanRepo *PengirimanRepo) UpdatePengiriman(data *models.Pengiriman, id string) error {
	err := pengirimanRepo.db.First(&models.Pengiriman{}, "id_pengiriman = ?", id).Error
	if err != nil {
		return err
	}

	err = pengirimanRepo.db.Model(&models.Pengiriman{}).Where("id_pengiriman = ?", id).Updates(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (pengirimanRepo *PengirimanRepo) DeletePengiriman(id int64) error {
	err := pengirimanRepo.db.First(&models.Pengiriman{}, "id_pengiriman = ?", id).Error
	if err != nil {
		return err
	}

	err = pengirimanRepo.db.Delete(&models.Pengiriman{}, pengirimanRepo.db.Where("id_pengiriman = ?", id)).Error
	if err != nil {
		return err
	}

	return nil
}
