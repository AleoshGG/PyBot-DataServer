package usecases

import (
	"PyBot-DataServer/work_periods/domain/models"
	"PyBot-DataServer/work_periods/domain/repositories"
)

type ReadingsRegister struct {
	wp repositories.IWorkPeriod
}

func NewReadingsRegister(wp repositories.IWorkPeriod) *ReadingsRegister {
	return &ReadingsRegister{wp: wp}
}

func (rr *ReadingsRegister) Run(r models.Reading) error {
	return rr.wp.ReadingsRegister(r)
}