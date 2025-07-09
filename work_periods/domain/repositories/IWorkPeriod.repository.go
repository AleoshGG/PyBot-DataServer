package repositories

import "PyBot-DataServer/work_periods/domain/models"

type IWorkPeriod interface {
	CreatePeriod(wp models.WorkPeriod) (error, period_id int)
	UpdatePeriod(end_hour string) (error) 
	ReadingsRegister(r models.Reading) (error)
	WasteCollectionRegister(wc models.WasteCollection) (error, id int)
	WeightRegister(w models.WeightData) (error, id int)
	GPSRegister(gps models.GPSData) (error, id int)
	GetWasteTypes() (error, []models.WasteType)
	CopyData()
}