package repository

import (
	"jasaPengiriman/models"

	"gorm.io/gorm"
)

type KurirRepo struct {
	db *gorm.DB
}

func NewKurirRepo(db *gorm.DB) *KurirRepo {
	return &KurirRepo{db}
}

func (kurirRepo *KurirRepo) GetAllKurir() []models.Kurir {
	var result []models.Kurir
	data := kurirRepo.db.Find(&result)

	if data.Error != nil {
		return nil
	}

	return result
}

func (kurirRepo *KurirRepo) CreateKurir(data *models.Kurir) error {
	err := kurirRepo.db.Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (kurirRepo *KurirRepo) UpdateKurir(data *models.Kurir, id string) error {
	err := kurirRepo.db.First(&models.Kurir{}, "id_kurir = ?", id).Error
	if err != nil {
		return err
	}

	err = kurirRepo.db.Model(&models.Kurir{}).Where("id_kurir = ?", id).Updates(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (kurirRepo *KurirRepo) DeleteKurir(id int64) error {
	var pengiriman []models.Pengiriman
	err := kurirRepo.db.First(&models.Kurir{}, "id_kurir = ?", id).Error
	if err != nil {
		return err
	}

	err = kurirRepo.db.Where("id_kurir = ?", id).Delete(&pengiriman).Error
	if err != nil {
		return err
	}

	err = kurirRepo.db.Delete(&models.Kurir{}, kurirRepo.db.Where("id_kurir = ?", id)).Error
	if err != nil {
		return err
	}

	return nil
}
