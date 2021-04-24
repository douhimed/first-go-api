package app

import (
	"productsapi/app/database"

	"github.com/gorilla/mux"
)


type App struct {
	Router *mux.Router
	DB database.ProductDb
}

func New() *App{
	
	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/api/products", a.CreateProductHandler()).Methods("POST")
}