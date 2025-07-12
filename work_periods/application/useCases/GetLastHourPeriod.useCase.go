package usecases

import (
	"PyBot-DataServer/work_periods/domain/models"
	"PyBot-DataServer/work_periods/domain/repositories"
)

type GetLastHourPeriod struct {
	wp repositories.IWorkPeriod
}

func NewGetLastHourPeriod(db repositories.IWorkPeriod) *GetLastHourPeriod {
	return &GetLastHourPeriod{wp: db}
}

func (glhp *GetLastHourPeriod) Run() (models.LastPeriod, error) {
	return glhp.wp.GetLastPeriod()
} 