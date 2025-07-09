package usecases

import (
	"PyBot-DataServer/sensors/domain/models"
	"PyBot-DataServer/sensors/domain/repositories"
)

type GetWasteTypes struct {
	sr repositories.ISensor
}

func NewGetWasteTypes(sr repositories.ISensor) *GetWasteTypes {
	return &GetWasteTypes{sr: sr}
}

func (gwt *GetWasteTypes) Run() ([]models.WasteType, error) {
	return gwt.sr.GetWasteTypes()
} 