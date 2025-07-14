package repositories

import "PyBot-DataServer/sensors/domain/models"

type ISensor interface {
	WasteCollectionRegister(wc models.WasteCollection) (int, error)
	UpdateWasteCollection(id int64) (error)
	WeightRegister(w models.WeightData) (int, error)
	GPSRegister(gps models.GPSData) (int, error)
	GetWasteTypes() ([]models.WasteType, error)
}