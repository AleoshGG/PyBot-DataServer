package controllers

import (
	"PyBot-DataServer/backup/application/services"
	"PyBot-DataServer/backup/infrastructure"
	"PyBot-DataServer/backup/infrastructure/adapters"
	"fmt"
	"log"
	"net/http"
	"reflect"

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
	ids, err := rbc.db.GetDataIdsBackupFalse()
	if err != nil {
		fmt.Printf("Error al obtener los ids de la tabla work_periods: %v", err)
	}
	dataTables, err := rbc.db.GetData()
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Error al obtener los datos: " + err.Error(),
		})
		return
	}

	rbc.serv.Run(dataTables)
	if len(dataTables) > 0 {
		// usa reflect para ver si Data es un slice y cuántos elementos tiene:
		v := reflect.ValueOf(dataTables[0].Data)
		if v.Kind() == reflect.Slice && v.Len() > 0 {
			// sólo si el slice tiene al menos 1 elemento, actualizamos:
			if err := rbc.db.UpdateIdsBackupDone(ids); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  false,
					"message": "Error al actualizar: " + err.Error(),
				})
				return
			}
		} else {
			// aquí puedes loguear o devolver otro status si quieres:
			log.Println("backup vacío, no se actualiza work_periods")
		}
	} else {
		log.Println("dataTables tiene 0 tablas")
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/sensors/",
		},
		"backup": dataTables,
	})
}