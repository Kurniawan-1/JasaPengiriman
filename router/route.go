package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	Handler "jasaPengiriman/jasaPengiriman/handler"
	Repo "jasaPengiriman/jasaPengiriman/repository"
	UC "jasaPengiriman/jasaPengiriman/usecase"
)

type Handlers struct {
	Ctx context.Context
	DB  *gorm.DB
	R   *gin.Engine
}

func (h *Handlers) Routes() {
	PelangganRepo := Repo.NewPelangganRepo(h.DB)
	PelangganUseCase := UC.NewPelangganUseCase(PelangganRepo)

	KurirRepo := Repo.NewKurirRepo(h.DB)
	KurirUseCase := UC.NewKurirUseCase(KurirRepo)

	BarangRepo := Repo.NewBarangRepo(h.DB)
	BarangUseCase := UC.NewBarangUseCase(BarangRepo)

	PengirimanRepo := Repo.NewPengirimanRepo(h.DB)
	PengirimanUseCase := UC.NewPengirimanUseCase(PengirimanRepo)

	v1 := h.R.Group("api")
	Handler.PelangganRoute(PelangganUseCase, v1)
	Handler.KurirRoute(KurirUseCase, v1)
	Handler.BarangRoute(BarangUseCase, v1)
	Handler.PengirimanRoute(PengirimanUseCase, v1)
}
