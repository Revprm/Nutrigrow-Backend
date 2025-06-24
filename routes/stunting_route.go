package routes

import (
	"github.com/Revprm/Nutrigrow-Backend/constants"
	"github.com/Revprm/Nutrigrow-Backend/controller"
	"github.com/Revprm/Nutrigrow-Backend/middleware"
	"github.com/Revprm/Nutrigrow-Backend/service"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Stunting(route *gin.Engine, injector *do.Injector) {
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	stuntingController := do.MustInvoke[controller.StuntingController](injector)

	routes := route.Group("/api/stunting")
	{
		routes.POST("/predict", stuntingController.Predict)

		authenticated := routes.Use(middleware.Authenticate(jwtService))
		{
			authenticated.POST("", stuntingController.Create)
			authenticated.GET("/:id", stuntingController.GetByID)
			authenticated.GET("/user/:user_id", stuntingController.GetByUserID)
			authenticated.GET("/user/:user_id/latest", stuntingController.GetLatestByUserID)
			authenticated.GET("/calendar", stuntingController.GetForCalendar)
			authenticated.PUT("/:id", stuntingController.Update)
			authenticated.DELETE("/:id", stuntingController.Delete)
		}
	}
}
