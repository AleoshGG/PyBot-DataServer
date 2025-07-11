package routes

import (
	"PyBot-DataServer/work_periods/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	workPeriodsRoutes := r.Group("/workPeriods")
	{
		workPeriodsRoutes.POST("/", controllers.NewCreatePeriodController().CreatePeriod)
		workPeriodsRoutes.GET("/", controllers.NewGetDistanceAndWeightController().GetDistanceAndWeight)
		workPeriodsRoutes.POST("/readings", controllers.NewReadingsRegisterController().ReadingsRegister)
		workPeriodsRoutes.PATCH("/", controllers.NewUpdatePeriodController().UpdatePeriod)
	}
}