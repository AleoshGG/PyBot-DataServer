package repositories

import "PyBot-DataServer/backup/domian/models"

type IService interface {
	SendDataTables(d []models.DataTable)
}