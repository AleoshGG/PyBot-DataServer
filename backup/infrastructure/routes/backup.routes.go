package routes

import (
	"PyBot-DataServer/backup/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	backupRouter := r.Group("/backup")
	{
		backupRouter.GET("/", controllers.NewRunBackupController().RunBackup)
	}
}