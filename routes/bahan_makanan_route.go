package routes

import (
	"github.com/Revprm/Nutrigrow-Backend/constants"
	"github.com/Revprm/Nutrigrow-Backend/controller"
	"github.com/Revprm/Nutrigrow-Backend/middleware"
	"github.com/Revprm/Nutrigrow-Backend/service"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func BahanMakanan(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	bahanMakananController := do.MustInvoke[controller.BahanMakananController](injector)

	routes := route.Group("/api/bahan-makanan")
	{
		routes.GET("/:id", bahanMakananController.GetByID)
		routes.GET("/nama/:nama", bahanMakananController.GetByNama)
		routes.GET("", bahanMakananController.GetAll)
		
		authenticated := routes.Use(middleware.Authenticate(jwtService))
		{
			authenticated.POST("", bahanMakananController.Create, middleware.OnlyAllow(constants.ENUM_ROLE_ADMIN))
			authenticated.PUT("/:id", bahanMakananController.Update, middleware.OnlyAllow(constants.ENUM_ROLE_ADMIN))
			authenticated.DELETE("/:id", bahanMakananController.Delete, middleware.OnlyAllow(constants.ENUM_ROLE_ADMIN))
		}
	}
}
