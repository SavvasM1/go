package model

import (
	"database/sql"
	"fmt"
	"log"
)

var connection *sql.DB

func Connect() *sql.DB {
	// dataSourceName := "root:1234@/mysql"
	// db, err := sql.Open("mysql", dataSourceName)
	// db, err := sql.Open("postgres", "mlpvysjgmtvtpm:2a66d1ad8c7e03adaa91991038aa12ce48b3018505773b9243bce36a019afecf@tcp(ec2-107-22-83-3.compute-1.amazonaws.com:5432)/d5ql5qe2lh0cko")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",

		"ec2-107-22-83-3.compute-1.amazonaws.com", 5432, "mlpvysjgmtvtpm", "2a66d1ad8c7e03adaa91991038aa12ce48b3018505773b9243bce36a019afecf", "d5ql5qe2lh0cko")

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Psql")
	connection = db
	return db
}
