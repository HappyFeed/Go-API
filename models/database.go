package models

import (
	"database/sql"
	"fmt"
	"log"

	"../config"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func CreateConnection() {
	url := config.GetUrlDatabase()
	if connection, err := sql.Open("mysql", url); err != nil {
		panic(err)
	} else {
		db = connection
	}
	//<username>:<password>@tcp(<host>:<port>)/<database>
}

func CloseConnection() {
	db.Close()
}

func createTable(tableName, schema string) {
	if !existsTable(tableName) {
		Exec(schema)
	} else {
		truncateTable(tableName)
	}
}

func CreateTables() {
	createTable("users", userSchema)
}

func existsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s' ", tableName)
	rows, _ := Query(sql)
	return rows.Next()
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println(err)
	}
	return rows, err
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func truncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}
