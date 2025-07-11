package infrastructure

import "PyBot-DataServer/sensors/infrastructure/adapters"

var postgres *adapters.PostgreSQL

func GoDependences() {
	postgres = adapters.NewPostgreSQL()
}

func GetPostgreSQL() *adapters.PostgreSQL {
	return postgres
}