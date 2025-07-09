package usecases

import "PyBot-DataServer/work_periods/domain/repositories"

type UpdatePeriod struct {
	wp repositories.IWorkPeriod
}

func NewUpdatePeriod(wp repositories.IWorkPeriod) *UpdatePeriod {
	return &UpdatePeriod{wp: wp}
}

func (up *UpdatePeriod) Run(end_hour string) (error) {
	return up.wp.UpdatePeriod(end_hour)
}