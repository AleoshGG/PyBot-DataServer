package usecases

import (
	"PyBot-DataServer/sensors/domain/models"
	"PyBot-DataServer/sensors/domain/repositories"
)

type WeightRegister struct {
	sr repositories.ISensor
}

func NewWeightRegister(sr repositories.ISensor) *WeightRegister {
	return &WeightRegister{sr: sr}
}

func (wr *WeightRegister) Run(w models.WeightData) (int, error) {
	return wr.sr.WeightRegister(w)
}