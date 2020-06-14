package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const username string = "root"
const password string = ""
const host string = "localhost"
const port int = 3306
const database string = "project_go_web"

func CreateConnection() {
	if connection, err := sql.Open("mysql", generateURL()); err != nil {
		panic(err)
	} else {
		db = connection
		fmt.Println("conexion exitosa")
	}
	//<username>:<password>@tcp(<host>:<port>)/<database>
}

func CloseConnection() {
	db.Close()
}

func generateURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
}
