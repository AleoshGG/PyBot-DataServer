package repositories

import "PyBot-DataServer/work_periods/domain/models"

type IWorkPeriod interface {
	GetLastPeriod() (models.LastPeriod, error)
	CreatePeriod(wp models.WorkPeriod) (int, error)
	UpdatePeriod(end_hour string, pediod_id int64) (error)
	GetDistanceAndWeight(period_id int64) (models.Reading, error) 
	ReadingsRegister(r models.Reading) (error)
}