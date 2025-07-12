package main

import (
	"PyBot-DataServer/database/conn"
	iS "PyBot-DataServer/sensors/infrastructure"
	rS "PyBot-DataServer/sensors/infrastructure/routes"
	iW "PyBot-DataServer/work_periods/infrastructure"
	rW "PyBot-DataServer/work_periods/infrastructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	//Migraciones 
	conn.Migration()

	// Cargar las dependencias
	iS.GoDependences()
	iW.GoDependences()

	// Inicio de la aplicación Gin
	r := gin.Default()

	// Configuración de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // o "*" para pruebas
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Rutas

	rS.RegisterRouter(r)
	rW.RegisterRouter(r)

	// Listen and Serve
	r.Run(":8080")
}