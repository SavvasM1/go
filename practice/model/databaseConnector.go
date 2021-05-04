package model

import (
	"database/sql"
	"fmt"
	"log"
)

var connection *sql.DB

func Connect() *sql.DB {
	dataSourceName := "root:1234@/mysql"
	db, err := sql.Open("mysql", dataSourceName)
	// db, err := sql.Open("mysql", "root:1234@(tcp:localhost:3306)")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MYSQL")
	connection = db
	return db
}
