package provider

import (
	"github.com/Revprm/Nutrigrow-Backend/config"
	"github.com/Revprm/Nutrigrow-Backend/constants"
	"github.com/Revprm/Nutrigrow-Backend/service"
	"github.com/samber/do"
	"gorm.io/gorm"
	"os"
)

func InitDatabase(injector *do.Injector) {
	do.ProvideNamed(injector, constants.DB, func(i *do.Injector) (*gorm.DB, error) {
		return config.SetUpDatabaseConnection(), nil
	})
}

func RegisterDependencies(injector *do.Injector) {
	InitDatabase(injector)

	do.ProvideNamed(injector, constants.JWTService, func(i *do.Injector) (service.JWTService, error) {
		return service.NewJWTService(), nil
	})

	// Initialize
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	// Get ML API URL from environment variable
	mlApiUrl := os.Getenv("ML_API_URL")
	if mlApiUrl == "" {
		// Default or error handling if not set
		mlApiUrl = "http://localhost:5000/predict" // Example default URL
	}

	// Provide Dependencies for all modules
	ProvideUserDependencies(injector, db, jwtService)
	ProvideBahanMakananDependencies(injector, db, jwtService)
	ProvideStuntingDependencies(injector, db, jwtService, mlApiUrl)
	ProvideBeritaDependencies(injector, db, jwtService)
	ProvideKategoriBeritaDependencies(injector, db, jwtService)
	ProvideMakananDependencies(injector, db, jwtService)
}
