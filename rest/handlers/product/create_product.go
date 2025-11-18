package product

import (
	"encoding/json"
	"net/http"
	"practice/database"
	"practice/utils"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "invalid json formate", 400)
		return
	}

	createProduct := database.Store(newProduct)

	utils.SendData(w, createProduct, 201)

}
