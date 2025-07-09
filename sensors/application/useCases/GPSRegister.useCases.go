package usecases

import (
	"PyBot-DataServer/sensors/domain/models"
	"PyBot-DataServer/sensors/domain/repositories"
)

type GPSRegister struct {
	sr repositories.ISensor
}

func NewGPSRegister(sr repositories.ISensor) *GPSRegister {
	return &GPSRegister{sr: sr}
}

func (gpsr *GPSRegister) Run(gps models.GPSData) (int, error) {
	return gpsr.sr.GPSRegister(gps)
}