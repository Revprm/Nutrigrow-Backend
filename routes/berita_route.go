package routes

import (
	"github.com/Revprm/Nutrigrow-Backend/constants"
	"github.com/Revprm/Nutrigrow-Backend/controller"
	"github.com/Revprm/Nutrigrow-Backend/middleware"
	"github.com/Revprm/Nutrigrow-Backend/service"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Berita(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	beritaController := do.MustInvoke[controller.BeritaController](injector)

	routes := route.Group("/api/berita")
	{
		routes.GET("/:id", beritaController.GetByID)
		routes.GET("/kategori/:kategori_id", beritaController.GetByKategori) 
		routes.GET("", beritaController.GetAllWithPagination) 

		authenticated := routes.Use(middleware.Authenticate(jwtService))
		{
			authenticated.POST("", beritaController.Create, middleware.OnlyAllow(constants.ENUM_ROLE_ADMIN))
			authenticated.PUT("/:id", beritaController.Update, middleware.OnlyAllow(constants.ENUM_ROLE_ADMIN))
			authenticated.DELETE("/:id", beritaController.Delete, middleware.OnlyAllow(constants.ENUM_ROLE_ADMIN))
		}
	}
}
