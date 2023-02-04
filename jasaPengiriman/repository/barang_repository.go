package repository

import (
	"jasaPengiriman/models"

	"gorm.io/gorm"
)

type BarangRepo struct {
	db *gorm.DB
}

func NewBarangRepo(db *gorm.DB) *BarangRepo {
	return &BarangRepo{db}
}

func (barangRepo *BarangRepo) GetAllBarang() []models.Barang {
	var result []models.Barang
	data := barangRepo.db.Find(&result)

	if data.Error != nil {
		return nil
	}

	return result
}

func (barangRepo *BarangRepo) CreateBarang(data *models.Barang) error {
	err := barangRepo.db.Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (barangRepo *BarangRepo) UpdateBarang(data *models.Barang, id string) error {
	err := barangRepo.db.First(&models.Barang{}, "id_barang = ?", id).Error
	if err != nil {
		return err
	}

	err = barangRepo.db.Model(&models.Barang{}).Where("id_barang = ?", id).Updates(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (barangRepo *BarangRepo) DeleteBarang(id int64) error {
	var pengiriman []models.Pengiriman
	err := barangRepo.db.First(&models.Barang{}, "id_barang = ?", id).Error
	if err != nil {
		return err
	}

	err = barangRepo.db.Where("id_barang = ?", id).Delete(&pengiriman).Error
	if err != nil {
		return err
	}
	err = barangRepo.db.Delete(&models.Barang{}, barangRepo.db.Where("id_barang = ?", id)).Error
	if err != nil {
		return err
	}

	return nil
}
