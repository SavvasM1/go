package model

import "fmt"

func DeleteEntry(name string) error {
	sqlStatement := "DELETE FROM TODO WHERE name=?"

	rows, err := connection.Query(sqlStatement, name)

	if err != nil {
		fmt.Println(err)
		return err
	}
	defer rows.Close()
	return nil
}
