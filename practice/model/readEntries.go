package model

import "github.com/SavvasM1/practice/views"

func ReadAll() ([]views.PostRequest, error) {
	rows, err := connection.Query("SELECT * FROM TODO")
	if err != nil {
		return nil, err
	}
	todos := []views.PostRequest{}

	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}
