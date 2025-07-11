package controllers

import (
	usecases "PyBot-DataServer/sensors/application/useCases"
	"PyBot-DataServer/sensors/domain/models"
	"PyBot-DataServer/sensors/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetWasteTypesController struct {
	gwt *usecases.GetWasteTypes
}

func NewGetWasteTypesController() *GetWasteTypesController {
	postgres := infrastructure.GetPostgreSQL()
	gwt := usecases.NewGetWasteTypes(postgres)

	return &GetWasteTypesController{gwt: gwt}
}

func (gwtc *GetWasteTypesController) GetWasteTypes(c *gin.Context) {
	var WasteTypes []models.WasteType
	
	WasteTypes, err := gwtc.gwt.Run()
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
			"self": "http://localhost:8080/sensors/",
		},
		"waste_types": WasteTypes,
	})
}