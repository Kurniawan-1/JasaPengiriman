package models

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Pelanggan struct {
	IDPelanggan  int    `gorm:"primaryKey;autoIncrement" json:"id_pelanggan"`
	Nama         string `json:"nama"`
	Alamat       string `json:"alamat"`
	NomorTelepon string `json:"nomor_telepon"`
}

type Kurir struct {
	IDKurir      int    `gorm:"primaryKey;autoIncrement" json:"id_kurir"`
	Nama         string `json:"nama"`
	Alamat       string `json:"alamat"`
	NomorTelepon string `json:"nomor_telepon"`
}

type Barang struct {
	IDBarang   int    `gorm:"primaryKey;autoIncrement" json:"id_barang"`
	NamaBarang string `json:"nama_barang"`
	Berat      int    `json:"berat"`
	Harga      int    `json:"harga"`
}

type Pengiriman struct {
	IDPengiriman      int       `gorm:"primaryKey;autoIncrement" json:"id_pengiriman"`
	TanggalPengiriman string    `json:"tanggal_pengiriman"`
	TanggalPenerimaan string    `json:"tanggal_penerimaan"`
	TanggalKirim      time.Time `json:"tanggal_kirim"`
	TanggalTerima     time.Time `json:"tanggal_terima"`
	AlamatTujuan      string    `json:"alamat_tujuan"`
	Status            string    `json:"status"`
	Biaya             int       `json:"biaya"`
	IdPelanggan       int       `json:"id_pelanggan"`
	IdKurir           int       `json:"id_kurir"`
	IdBarang          int       `json:"id_barang"`
	Pelanggan         Pelanggan `gorm:"foreignKey:IdPelanggan;references:IDPelanggan"`
	Kurir             Kurir     `gorm:"foreignKey:IdKurir;references:IDKurir"`
	Barang            Barang    `gorm:"foreignKey:IdBarang;references:IDBarang"`
}

// //////////////////////////////////////////////////////////////////////////////////
type PelangganRepo interface {
	GetAllPelanggan() []Pelanggan
	CreatePelanggan(*Pelanggan) error
	UpdatePelanggan(p *Pelanggan, id string) error
	DeletePelanggan(id int64) error
}

type PelangganUseCase interface {
	GetAllPelanggan() []Pelanggan
	CreatePelanggan(*gin.Context) error
	UpdatePelanggan(*gin.Context) error
	DeletePelanggan(*gin.Context) error
}

// /////////////////////////////////////////////////////////////////////////////////
type KurirRepo interface {
	GetAllKurir() []Kurir
	CreateKurir(*Kurir) error
	UpdateKurir(k *Kurir, id string) error
	DeleteKurir(id int64) error
}

type KurirUseCase interface {
	GetAllKurir() []Kurir
	CreateKurir(*gin.Context) error
	UpdateKurir(*gin.Context) error
	DeleteKurir(*gin.Context) error
}

// /////////////////////////////////////////////////////////////////////////////////
type BarangRepo interface {
	GetAllBarang() []Barang
	CreateBarang(*Barang) error
	UpdateBarang(b *Barang, id string) error
	DeleteBarang(id int64) error
}

type BarangUseCase interface {
	GetAllBarang() []Barang
	CreateBarang(*gin.Context) error
	UpdateBarang(*gin.Context) error
	DeleteBarang(*gin.Context) error
}

// /////////////////////////////////////////////////////////////////////////////////
type PengirimanRepo interface {
	GetAllPengiriman() []Pengiriman
	CreatePengiriman(*Pengiriman) error
	UpdatePengiriman(p *Pengiriman, id string) error
	DeletePengiriman(id int64) error
}

type PengirimanUseCase interface {
	GetAllPengiriman() []Pengiriman
	CreatePengiriman(*gin.Context) error
	UpdatePengiriman(*gin.Context) error
	DeletePengiriman(*gin.Context) error
}
