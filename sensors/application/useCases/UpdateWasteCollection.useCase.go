package usecases

import "PyBot-DataServer/sensors/domain/repositories"

type UpdateWasteCollection struct {
	sr repositories.ISensor
}

func NewUpdateWasteCollection(db repositories.ISensor) *UpdateWasteCollection {
	return &UpdateWasteCollection{sr: db}
}

func (upc *UpdateWasteCollection) Run(id int64) (error) {
	return upc.sr.UpdateWasteCollection(id)
}