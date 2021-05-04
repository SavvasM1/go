package model

import "github.com/SavvasM1/practice/views"

func ReadByName(name string) ([]views.PostRequest, error) {
	sqlStatement := "SELECT * FROM TODO WHERE name=?"

	rows, err := connection.Query(sqlStatement, name)

	if err != nil {
		return nil, err
	}

	entries := []views.PostRequest{}

	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		entries = append(entries, data)
	}

	return entries, nil
}
