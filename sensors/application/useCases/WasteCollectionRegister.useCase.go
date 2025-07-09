package usecases

import (
	"PyBot-DataServer/sensors/domain/models"
	"PyBot-DataServer/sensors/domain/repositories"
)

type WasteCollectionRegister struct {
	sr repositories.ISensor
}

func NewWasteCollectionRegister(sr repositories.ISensor) *WasteCollectionRegister {
	return &WasteCollectionRegister{sr: sr}
}

func (wcr *WasteCollectionRegister) Run(wc models.WasteCollection) (int, error) {
	return wcr.sr.WasteCollectionRegister(wc)
}