package provider

import (
	// "github.com/Revprm/Nutrigrow-Backend/constants"
	"github.com/Revprm/Nutrigrow-Backend/controller"
	"github.com/Revprm/Nutrigrow-Backend/repository"
	"github.com/Revprm/Nutrigrow-Backend/service"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideBeritaDependencies(injector *do.Injector, db *gorm.DB, jwtService service.JWTService) {
	// Repository
	beritaRepository := repository.NewBeritaRepository(db)
	kategoriBeritaRepository := repository.NewKategoriBeritaRepository(db)
	// Service
	beritaService := service.NewBeritaService(beritaRepository, kategoriBeritaRepository, db)

	// Controller
	do.Provide(
		injector, func(i *do.Injector) (controller.BeritaController, error) {
			return controller.NewBeritaController(beritaService), nil
		},
	)
}
