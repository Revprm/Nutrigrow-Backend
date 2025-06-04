package provider

import (
	// "github.com/Revprm/Nutrigrow-Backend/constants"
	"github.com/Revprm/Nutrigrow-Backend/controller"
	"github.com/Revprm/Nutrigrow-Backend/repository"
	"github.com/Revprm/Nutrigrow-Backend/service"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideBahanMakananDependencies(injector *do.Injector, db *gorm.DB, jwtService service.JWTService) {
	// Repository
	bahanMakananRepository := repository.NewBahanMakananRepository(db)

	// Service
	bahanMakananService := service.NewBahanMakananService(bahanMakananRepository, db)

	// Controller
	do.Provide(
		injector, func(i *do.Injector) (controller.BahanMakananController, error) {
			return controller.NewBahanMakananController(bahanMakananService), nil
		},
	)
}
