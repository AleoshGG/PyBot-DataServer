package controllers

import (
	usecases "PyBot-DataServer/sensors/application/useCases"
	"PyBot-DataServer/sensors/domain/models"
	"PyBot-DataServer/sensors/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GPSRegisterController struct {
	gpsr *usecases.GPSRegister
}

func NewGPSRegisterController() *GPSRegisterController {
	postgres := infrastructure.GetPostgreSQL()
	gpsr := usecases.NewGPSRegister(postgres)

	return &GPSRegisterController{gpsr: gpsr}
}

func (gpsrc *GPSRegisterController) GPSRegister(c *gin.Context) {
	var register models.GPSData

	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Error en cuerpo del mensaje: " + err.Error(),
		})
		return
	}

	id, err := gpsrc.gpsr.Run(register)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Error al insertar el recurso: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data": gin.H{
			"type": "GPS_data",
			"gps_data_id": id,
			/*"attributes": gin.H{
				"first_name": user.First_name,
				"last_name": user.Last_name,
				"email": user.Email,
				"password": user.Password,
			,}*/
		},
	})
}