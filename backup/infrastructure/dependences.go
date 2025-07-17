package infrastructure

import (
	"PyBot-DataServer/backup/infrastructure/adapters"
	"fmt"
)

var postgre *adapters.PostgreSQL
var rabbit *adapters.RabbitMQ

func GoDependences() {
	postgre = adapters.NewPostgreSQL()
	r, err := adapters.NewRabbitMQ()
	if err != nil {
		fmt.Printf("Error al instanciar: %v", err)
	}
	rabbit = r
}

func GetPostgreSQL() *adapters.PostgreSQL {
	return postgre
}

func GetRabbitMQ() *adapters.RabbitMQ {
	return rabbit
}
