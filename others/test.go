package main

import (
	"./models"
)

func main() {
	models.CreateConnection()
	models.Ping()
}
