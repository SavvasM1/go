package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SavvasM1/practice/model"
	"github.com/SavvasM1/practice/views"
)

func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			//New object PostRequest
			data := views.PostRequest{}

			// Deserializing json body
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)

			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				fmt.Println(err.Error())
				w.Write([]byte("Error occured while trying to write"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			if r.URL.Query().Get("name") != "" {
				fmt.Println("Found query param: " + r.URL.Query().Get("name"))
				nameQueryString := r.URL.Query().Get("name")
				data, err := model.ReadByName(nameQueryString)
				if err != nil {
					w.Write([]byte(err.Error()))
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(data)
			} else {
				fmt.Println("No query param passed")
				data, err := model.ReadAll()
				if err != nil {
					w.Write([]byte(err.Error()))
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(data)
			}
		} else if r.Method == http.MethodDelete {
			name := r.URL.Path[1:]
			if err := model.DeleteEntry(name); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Status string `json:status`
			}{"Item deleted"})
		}
	}
}
