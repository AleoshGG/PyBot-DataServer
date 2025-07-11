package controllers

import (
	usecases "PyBot-DataServer/sensors/application/useCases"
	"PyBot-DataServer/sensors/domain/models"
	"PyBot-DataServer/sensors/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WasteCollectionRegisterController struct {
	wcr *usecases.WasteCollectionRegister
}

func NewWasteCollectionRegisterController() *WasteCollectionRegisterController {
	postgres := infrastructure.GetPostgreSQL()
	wcr := usecases.NewWasteCollectionRegister(postgres)

	return &WasteCollectionRegisterController{wcr: wcr}
}

func (wcrc *WasteCollectionRegisterController) WasteCollectionRegister(c *gin.Context) {
	var register models.WasteCollection

	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Error en cuerpo del mensaje: " + err.Error(),
		})
		return
	}

	id, err := wcrc.wcr.Run(register)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Error al insertar el recurso: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data": gin.H{
			"type": "waste_collection",
			"waste_collection_id": id,
			"attributes": gin.H{
				"period_id": register.Period_id,
				"amount": register.Amount,
				"waste_id": register.Waste_id,
			},
		},
	})
}