package controllers

import (
	usecases "PyBot-DataServer/sensors/application/useCases"
	"PyBot-DataServer/sensors/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateWasteCollectionContoller struct {
	uwc *usecases.UpdateWasteCollection
}

func NewUpdateWasteCollectionController() *UpdateWasteCollectionContoller {
	postgres := infrastructure.GetPostgreSQL()
	uwc := usecases.NewUpdateWasteCollection(postgres)

	return &UpdateWasteCollectionContoller{uwc: uwc}
}

func (uwcc *UpdateWasteCollectionContoller) UpdateWasteCollection(c *gin.Context) {
	Id := c.Query("Id")

	if Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Error: Faltan los parámetros 'Amount' e 'Id'",
		})
		return
	}

	id, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Error al obtener 'Id': " + err.Error(),
		})
		return
	}

	if err := uwcc.uwc.Run(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "Error al actualizar el recurso: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Recurso actualizado",
	})
}