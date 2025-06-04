package provider

import (
	// "github.com/Revprm/Nutrigrow-Backend/constants"
	"github.com/Revprm/Nutrigrow-Backend/controller"
	"github.com/Revprm/Nutrigrow-Backend/repository"
	"github.com/Revprm/Nutrigrow-Backend/service"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideKategoriBeritaDependencies(injector *do.Injector, db *gorm.DB, jwtService service.JWTService) {
	// Repository
	kategoriBeritaRepository := repository.NewKategoriBeritaRepository(db)

	// Service
	kategoriBeritaService := service.NewKategoriBeritaService(kategoriBeritaRepository, db)

	// Controller
	do.Provide(
		injector, func(i *do.Injector) (controller.KategoriBeritaController, error) {
			return controller.NewKategoriBeritaController(kategoriBeritaService), nil
		},
	)
}
