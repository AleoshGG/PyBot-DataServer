package controllers

import (
	usecases "PyBot-DataServer/work_periods/application/useCases"
	"PyBot-DataServer/work_periods/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetLastHourPeriodController struct {
	glhp *usecases.GetLastHourPeriod
}

func NewGetLastHourPeriodController() *GetLastHourPeriodController {
	postgres := infrastructure.GetPostgreSQL()
	glhp := usecases.NewGetLastHourPeriod(postgres)

	return &GetLastHourPeriodController{glhp: glhp}
}

func (glhpc *GetLastHourPeriodController) GetLastPeriod(c *gin.Context) {
	
	lastPeriod, err := glhpc.glhp.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Error al obtener el recurso: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/work_periods/",
		},
		"last_period": lastPeriod,
	})

}