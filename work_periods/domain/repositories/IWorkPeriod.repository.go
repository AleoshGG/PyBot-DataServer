package repositories

import "PyBot-DataServer/work_periods/domain/models"

type IWorkPeriod interface {
	CreatePeriod(wp models.WorkPeriod) (error, period_id int)
	UpdatePeriod(end_hour string) (error)
	GetDistanceAndWeight() (error, models.Reading) 
	ReadingsRegister(r models.Reading) (error)
}