package controllers

import (
	usecases "PyBot-DataServer/work_periods/application/useCases"
	"PyBot-DataServer/work_periods/domain/models"
	"PyBot-DataServer/work_periods/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetDistanceAndWeightController struct {
	gdw *usecases.GetDistanceAndWeight
}

func NewGetDistanceAndWeightController() *GetDistanceAndWeightController {
	postgres := infrastructure.GetPostgreSQL()
	gdw := usecases.NewGetDistanceAndWeight(postgres)
	return &GetDistanceAndWeightController{gdw: gdw}
}

func (gdwc *GetDistanceAndWeightController) GetDistanceAndWeight(c *gin.Context) {
	id := c.Query("id")
	
	var reading models.Reading
	
	period_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Error: No se pudo obtener el ID numérico",
		})
		return
	}

	reading, err = gdwc.gdw.Run(period_id)
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
		"last_reading": reading,
	})
}