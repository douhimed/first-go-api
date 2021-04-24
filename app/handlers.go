package app

import (
	"fmt"
	"log"
	"net/http"
	"productsapi/app/models"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintf(w, "Welcome to Products API")
	}
}

func (a *App) CreateProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		req := models.ProductRequest{}
		
		err := parse(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse the body, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		p := &models.Product{
			ID: 0,
			Label: req.Label,
			Quantity: req.Quantity,
			Price: req.Price,
		}

		err = a.DB.SaveProduct(p)
		if err != nil {
			log.Printf("Cannot save the product, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := mapProductToJson(p)
		sendResponse(w, r, resp, http.StatusOK)
	}
}