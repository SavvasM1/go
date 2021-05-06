package model

func CreateTodo(name, todo string) error {
	insertQ, err := connection.Query("INSERT INTO TODO VALUES ($1, $2)", name, todo)
	if err != nil {
		return err
	}

	defer insertQ.Close()

	return nil
}
