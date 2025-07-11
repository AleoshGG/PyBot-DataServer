package controllers

import (
	usecases "PyBot-DataServer/sensors/application/useCases"
	"PyBot-DataServer/sensors/domain/models"
	"PyBot-DataServer/sensors/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WeightRegisterController struct {
	wr *usecases.WeightRegister
}

func NewWeightRegisterController() *WeightRegisterController {
	postgres := infrastructure.GetPostgreSQL()
	wr := usecases.NewWeightRegister(postgres)

	return &WeightRegisterController{wr: wr}
}

func (wrc *WeightRegisterController) WeightRegister(c *gin.Context) {
	var register models.WeightData

	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Error en cuerpo del mensaje: " + err.Error(),
		})
		return
	}

	id, err := wrc.wr.Run(register)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Error al insetar el recurso: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data": gin.H{
			"type": "weight_data",
			"weight_data_id": id,
			"attributes": gin.H{
				"period_id": register.Period_id,
				"hour_period": register.Hour_period,
				"weight": register.Weight,
			},
		},
	})
}