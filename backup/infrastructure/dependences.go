package infrastructure

import "PyBot-DataServer/backup/infrastructure/adapters"

var postgre *adapters.PostgreSQL
var rabbit *adapters.RabbitMQ

func GoDependences() {
	postgre = adapters.NewPostgreSQL()
	rabbit = adapters.NewRabbitMQ()
}

func GetPostgreSQL() *adapters.PostgreSQL {
	return postgre
}

func GetRabbitMQ() *adapters.RabbitMQ {
	return rabbit
}
