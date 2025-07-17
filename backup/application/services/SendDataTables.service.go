package services

import (
	"PyBot-DataServer/backup/application/repositories"
	"PyBot-DataServer/backup/domian/models"
)

type SendDataTables struct {
	s repositories.IService
}

func NewSendDataTablesService(s repositories.IService) *SendDataTables {
	return &SendDataTables{s: s}
}

func (serv *SendDataTables) Run(d []models.DataTable) {
	serv.s.SendDataTables(d)
}