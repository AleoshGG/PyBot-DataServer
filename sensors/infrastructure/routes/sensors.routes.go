package routes

import (
	"PyBot-DataServer/sensors/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	sensorsRouter := r.Group("/sensors")
	{
		sensorsRouter.GET("/", controllers.NewGetWasteTypesController().GetWasteTypes)
		sensorsRouter.POST("/gps", controllers.NewGPSRegisterController().GPSRegister)
		sensorsRouter.POST("/waste", controllers.NewWasteCollectionRegisterController().WasteCollectionRegister)
		sensorsRouter.POST("/weight", controllers.NewWeightRegisterController().WeightRegister)
		sensorsRouter.PATCH("/", controllers.NewUpdateWasteCollectionController().UpdateWasteCollection)
	}
}