package controllers

import (
	"PyBot-DataServer/backup/application/services"
	"PyBot-DataServer/backup/infrastructure"
	"PyBot-DataServer/backup/infrastructure/adapters"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RunBackupController struct {
	db *adapters.PostgreSQL
	serv *services.SendDataTables
}

func NewRunBackupController() *RunBackupController {
	db := infrastructure.GetPostgreSQL()
	rabbit := infrastructure.GetRabbitMQ()
	service := services.NewSendDataTablesService(rabbit)
	return &RunBackupController{db: db, serv: service}
}

func (rbc *RunBackupController) RunBackup(c *gin.Context) {
	dataTables, err := rbc.db.GetData()
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Error al obtener los datos: " + err.Error(),
		})
		return
	}

	rbc.serv.Run(dataTables)
	if len(dataTables) < 1 {
		
		err = rbc.db.UpdateIdsBackupDone()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "Error al actualizar: " + err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/sensors/",
		},
		"backup": dataTables,
	})
}