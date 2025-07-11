package infrastructure

import "PyBot-DataServer/work_periods/infrastructure/adapters"

var postgre *adapters.PostgreSQL

func GoDependences() {
	postgre = adapters.NewPostgreSQL()
}

func GetPostgreSQL() *adapters.PostgreSQL {
	return postgre
}