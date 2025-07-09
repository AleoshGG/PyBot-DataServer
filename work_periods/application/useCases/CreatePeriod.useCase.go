package usecases

import (
	"PyBot-DataServer/work_periods/domain/models"
	"PyBot-DataServer/work_periods/domain/repositories"
)

type CreatePeriod struct {
	wp repositories.IWorkPeriod
}

func NewCreatePeriod(wp repositories.IWorkPeriod) *CreatePeriod{
	return &CreatePeriod{wp: wp}
}

func (cp *CreatePeriod) Run(wp models.WorkPeriod) (error, int) {
	return cp.wp.CreatePeriod(wp)
}