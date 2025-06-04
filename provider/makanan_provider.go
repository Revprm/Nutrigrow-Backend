package provider

import (
	// "github.com/Revprm/Nutrigrow-Backend/constants"
	"github.com/Revprm/Nutrigrow-Backend/controller"
	"github.com/Revprm/Nutrigrow-Backend/repository"
	"github.com/Revprm/Nutrigrow-Backend/service"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideMakananDependencies(injector *do.Injector, db *gorm.DB, jwtService service.JWTService) {
	// Repository
	makananRepository := repository.NewMakananRepository(db)
	bahanMakananRepository := repository.NewBahanMakananRepository(db)

	// Service
	makananService := service.NewMakananService(makananRepository, bahanMakananRepository, db)

	// Controller
	do.Provide(
		injector, func(i *do.Injector) (controller.MakananController, error) {
			return controller.NewMakananController(makananService), nil
		},
	)
}
