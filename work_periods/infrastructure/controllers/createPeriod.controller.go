package controllers

import (
	usecases "PyBot-DataServer/work_periods/application/useCases"
	"PyBot-DataServer/work_periods/domain/models"
	"PyBot-DataServer/work_periods/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePeriodController struct {
	cp *usecases.CreatePeriod
}

func NewCreatePeriodController() *CreatePeriodController {
	postgre := infrastructure.GetPostgreSQL()
	cp := usecases.NewCreatePeriod(postgre)
	return &CreatePeriodController{cp: cp}
}

func (cpc *CreatePeriodController) CreatePeriod(c *gin.Context) {
	var period models.WorkPeriod

    if err := c.ShouldBindJSON(&period); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}

	period_id, err := cpc.cp.Run(period)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "No se pudo guardar el recurso " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data": gin.H{
			"type": "work_periods",
			"work_periods_id": period_id,
			"attributes": gin.H{
				"start_hour": period.Start_hour,
				"end_hour": period.End_hour,
				"day_work": period.Day_work,
				"prototype_id": period.Prototype_id,
			},
		},
	})
}