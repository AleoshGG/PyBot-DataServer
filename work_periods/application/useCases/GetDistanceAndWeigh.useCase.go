package usecases

import (
	"PyBot-DataServer/work_periods/domain/models"
	"PyBot-DataServer/work_periods/domain/repositories"
)

type GetDistanceAndWeight struct {
	wp repositories.IWorkPeriod
}

func NewGetDistanceAndWeight(wp repositories.IWorkPeriod) *GetDistanceAndWeight {
	return &GetDistanceAndWeight{wp: wp}
}

func (gdw *GetDistanceAndWeight) Run() (models.Reading, error) {
	return gdw.wp.GetDistanceAndWeight()
}