package repositories

import "PyBot-DataServer/sensors/domain/models"

type ISensor interface {
	WasteCollectionRegister(wc models.WasteCollection) (error, id int)
	WeightRegister(w models.WeightData) (error, id int)
	GPSRegister(gps models.GPSData) (error, id int)
	GetWasteTypes() (error, []models.WasteType)
}