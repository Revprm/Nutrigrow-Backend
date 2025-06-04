package routes

import (
	"github.com/Revprm/Nutrigrow-Backend/constants"
	"github.com/Revprm/Nutrigrow-Backend/controller"
	"github.com/Revprm/Nutrigrow-Backend/middleware"
	"github.com/Revprm/Nutrigrow-Backend/service"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Makanan(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	makananController := do.MustInvoke[controller.MakananController](injector)

	routes := route.Group("/api/makanan")
	{
		routes.GET("/:id", makananController.GetByID)
		routes.GET("/nama/:nama", makananController.GetByNama)
		routes.GET("", makananController.GetAll) // Get all food items with pagination
		routes.GET("/bahan/:bahanId", makananController.GetByBahanMakanan) // Get food items by ingredient ID with pagination

		authenticated := routes.Use(middleware.Authenticate(jwtService))
		{
			authenticated.POST("", makananController.Create)
			authenticated.PUT("/:id", makananController.Update)
			authenticated.DELETE("/:id", makananController.Delete)
		}
	}
}
