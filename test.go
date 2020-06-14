package main

import (
	"fmt"

	"./models"
)

func main() {
	models.CreateConnection()
	models.CreateTables()

	models.CreateUser("Alejandro", "1234", "asldl@gmail.com")
	models.CreateUser("Juan", "1234", "asldl@gmail.com")
	models.CreateUser("Camilo", "1234", "asldl@gmail.com")

	users := models.GetUsers()
	fmt.Println(users)

	models.CloseConnection()

}
