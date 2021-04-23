package main

import (
	"log"
	"net/http"
	"os"
	"productsapi/app"
	"productsapi/app/database"

	_ "github.com/go-sql-driver/mysql"
)


func main() {

	app := app.New()
	app.DB = &database.DB{}
	err := app.DB.Open()
	check(err)

	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)
	 
	err = http.ListenAndServe(":9090", nil)
	check(err)

}

func check(e error){
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}