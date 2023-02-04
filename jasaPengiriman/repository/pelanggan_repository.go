package repository

import (
	"jasaPengiriman/models"

	"gorm.io/gorm"
)

type PelangganRepo struct {
	db *gorm.DB
}

func NewPelangganRepo(db *gorm.DB) *PelangganRepo {
	return &PelangganRepo{db}
}

func (pelangganRepo *PelangganRepo) GetAllPelanggan() []models.Pelanggan {
	var result []models.Pelanggan
	data := pelangganRepo.db.Find(&result)

	if data.Error != nil {
		return nil
	}

	return result
}

func (pelangganRepo *PelangganRepo) CreatePelanggan(data *models.Pelanggan) error {
	err := pelangganRepo.db.Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (pelangganRepo *PelangganRepo) UpdatePelanggan(data *models.Pelanggan, id string) error {
	err := pelangganRepo.db.First(&models.Pelanggan{}, "id_pelanggan = ?", id).Error
	if err != nil {
		return err
	}

	err = pelangganRepo.db.Model(&models.Pelanggan{}).Where("id_pelanggan = ?", id).Updates(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (pelangganRepo *PelangganRepo) DeletePelanggan(id int64) error {
	var pengiriman []models.Pengiriman
	err := pelangganRepo.db.First(&models.Pelanggan{}, "id_pelanggan = ?", id).Error
	if err != nil {
		return err
	}

	err = pelangganRepo.db.Where("id_pelanggan = ?", id).Delete(&pengiriman).Error
	if err != nil {
		return err
	}

	err = pelangganRepo.db.Delete(&models.Pelanggan{}, pelangganRepo.db.Where("id_pelanggan = ?", id)).Error
	if err != nil {
		return err
	}

	return nil
}
