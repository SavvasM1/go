package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SavvasM1/practice/controller"
	"github.com/SavvasM1/practice/model"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	mux := controller.Register()

	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving ...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}

// TODO: Make db conn an interface:
// https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/
