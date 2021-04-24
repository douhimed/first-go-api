package app

import (
	"encoding/json"
	"log"
	"net/http"
	"productsapi/app/models"
)


func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func sendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json, err=%v\n", err)
	}
}

func mapProductToJson(p *models.Product) models.ProductResponse {
	return models.ProductResponse{
		ID: p.ID,
		Label: p.Label,
		Quantity: p.Quantity,
		Price: p.Price,
	}
}