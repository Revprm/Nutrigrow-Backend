package provider

import (
	// "github.com/Revprm/Nutrigrow-Backend/constants"
	"github.com/Revprm/Nutrigrow-Backend/controller"
	"github.com/Revprm/Nutrigrow-Backend/repository"
	"github.com/Revprm/Nutrigrow-Backend/service"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func ProvideStuntingDependencies(injector *do.Injector, db *gorm.DB, jwtService service.JWTService, mlApiUrl string) {
	// Repository
	stuntingRepository := repository.NewStuntingRepository(db)
	
	// Service
	stuntingService := service.NewStuntingService(stuntingRepository, db, mlApiUrl)

	// Controller
	do.Provide(
		injector, func(i *do.Injector) (controller.StuntingController, error) {
			return controller.NewStuntingController(stuntingService), nil
		},
	)
}
