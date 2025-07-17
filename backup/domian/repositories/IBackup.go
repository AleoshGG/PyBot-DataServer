package repositories

import "PyBot-DataServer/backup/domian/models"

type IBackup interface {
	GetDataIdsBackupFalse() ([]models.DataTable, error)
	UpdateIdsBackupDone([]string) error
}