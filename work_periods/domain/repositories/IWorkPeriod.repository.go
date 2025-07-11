package repositories

import "PyBot-DataServer/work_periods/domain/models"

type IWorkPeriod interface {
	CreatePeriod(wp models.WorkPeriod) (int, error)
	UpdatePeriod(end_hour string, pediod_id int64) (error)
	GetDistanceAndWeight() ( models.Reading, error) 
	ReadingsRegister(r models.Reading) (error)
}