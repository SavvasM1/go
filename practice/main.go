package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SavvasM1/practice/controller"
	"github.com/SavvasM1/practice/model"

	_ "github.com/lib/pq"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT not set")
		return ":3000", nil
	}
	return ":" + port, nil
}

func main() {
	addr, err := determineListenAddress()

	if err != nil {
		log.Fatal(err)
	}
	mux := controller.Register()

	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving ...")
	log.Fatal(http.ListenAndServe(addr, mux))
}
