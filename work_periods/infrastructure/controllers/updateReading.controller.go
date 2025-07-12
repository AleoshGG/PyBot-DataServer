package controllers

import (
	usecases "PyBot-DataServer/work_periods/application/useCases"
	"PyBot-DataServer/work_periods/domain/models"
	"PyBot-DataServer/work_periods/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateReadingController struct {
	ur usecases.UpdateReading
}

func NewUpdateReadingController() *UpdateReadingController {
	postgre := infrastructure.GetPostgreSQL()
	ur := usecases.NewUpdateReadin(postgre)

	return &UpdateReadingController{ur: *ur}
}

func (urc *UpdateReadingController) UpdateReading(c *gin.Context) {
	var reading models.Reading

	if err := c.ShouldBindJSON(&reading); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Error en cuerpo del mensaje: " + err.Error(),
		})
		return
	}

	if err := urc.ur.Run(reading); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "Error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Recurso actualizado",
	})	
}