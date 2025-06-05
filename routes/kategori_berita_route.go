package routes

import (
	"github.com/Revprm/Nutrigrow-Backend/constants"
	"github.com/Revprm/Nutrigrow-Backend/controller"
	"github.com/Revprm/Nutrigrow-Backend/middleware"
	"github.com/Revprm/Nutrigrow-Backend/service"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func KategoriBerita(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	kategoriBeritaController := do.MustInvoke[controller.KategoriBeritaController](injector)

	routes := route.Group("/api/kategori-berita")
	{
		routes.GET("/:id", kategoriBeritaController.GetByID)
		routes.GET("/nama", kategoriBeritaController.GetByNama)
		routes.GET("", kategoriBeritaController.GetAll)

		authenticated := routes.Use(middleware.Authenticate(jwtService))
		{
			authenticated.POST("", kategoriBeritaController.Create, middleware.OnlyAllow(constants.ENUM_ROLE_ADMIN))
			authenticated.PUT("/:id", kategoriBeritaController.Update, middleware.OnlyAllow(constants.ENUM_ROLE_ADMIN))
			authenticated.DELETE("/:id", kategoriBeritaController.Delete, middleware.OnlyAllow(constants.ENUM_ROLE_ADMIN))
		}
	}
}
