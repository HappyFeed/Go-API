package main

import (
	"fmt"

	"./orm"
)

func main() {
	orm.CreateConnection()
	orm.CreateTables()

	user := orm.NewUser("Alejandro", "123", "asldkas@gmail.com")
	user.Save()

	users := orm.GetUsers()
	fmt.Println(users)

	user = orm.GetUser(1)
	user.Username = "Andres"
	user.Save()
	fmt.Println(user)

	user.Delete()

	orm.CloseConnection()
}
