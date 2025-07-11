package controllers

import (
	usecases "PyBot-DataServer/work_periods/application/useCases"
	"PyBot-DataServer/work_periods/domain/models"
	"PyBot-DataServer/work_periods/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReadingsRegisterController struct {
	rr usecases.ReadingsRegister
}

func NewReadingsRegisterController() *ReadingsRegisterController {
	postgre := infrastructure.GetPostgreSQL()
	rr := usecases.NewReadingsRegister(postgre)

	return &ReadingsRegisterController{rr: *rr}
}

func (rrc *ReadingsRegisterController) ReadingsRegister(c *gin.Context) {
	var reading models.Reading

	if err := c.ShouldBindJSON(&reading); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Error en cuerpo del mensaje: " + err.Error(),
		})
		return
	}

	if err := rrc.rr.Run(reading); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Error al insertar: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data": gin.H{
			"type": "readings",
			"period_id": reading.Period_id,
			"attributes": gin.H{
				"distance_traveled": reading.Distance_traveled,
				"weight_waste": reading.Weight_waste,
			},
		},
	})
}