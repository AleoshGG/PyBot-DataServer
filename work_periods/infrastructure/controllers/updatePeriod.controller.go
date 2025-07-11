package controllers

import (
	usecases "PyBot-DataServer/work_periods/application/useCases"
	"PyBot-DataServer/work_periods/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdatePeriodController struct {
	up *usecases.UpdatePeriod
}

func NewUpdatePeriodController() *UpdatePeriodController {
	postgre := infrastructure.GetPostgreSQL()
	up := usecases.NewUpdatePeriod(postgre)

	return &UpdatePeriodController{up: up}
}

func (upc *UpdatePeriodController) UpdatePeriod(c *gin.Context) {
    endHour := c.Query("endHour")
	id := c.Query("id")

	if endHour == "" || id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Error: endHour y id son requeridos",
		})
		return
	}

	period_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Error: No se pudo obtener el ID numérico",
		})
		return
	}

	if err := upc.up.Run(endHour, period_id); err != nil {
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
