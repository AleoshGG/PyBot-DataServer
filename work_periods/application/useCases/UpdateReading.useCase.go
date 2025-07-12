package usecases

import (
	"PyBot-DataServer/work_periods/domain/models"
	"PyBot-DataServer/work_periods/domain/repositories"
)

type UpdateReading struct {
	wr repositories.IWorkPeriod
}

func NewUpdateReadin(db repositories.IWorkPeriod) *UpdateReading {
	return &UpdateReading{wr: db}
}

func (ur *UpdateReading) Run(r models.Reading) (error) {
	return ur.wr.UpdateReadings(r)
}