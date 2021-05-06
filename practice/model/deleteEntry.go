package model

import "fmt"

func DeleteEntry(name string) error {
	sqlStatement := "DELETE FROM TODO WHERE name=$1"

	fmt.Println("DELETE STATEMENT = " + sqlStatement)
	fmt.Println("NAME PROVIDED = " + name)
	rows, err := connection.Query(sqlStatement, name)

	if err != nil {
		fmt.Println(err)
		return err
	}
	defer rows.Close()
	return nil
}
