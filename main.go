package main

import (
	"PyBot-DataServer/database/conn"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	conn.Migration()
}