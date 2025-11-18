package product

import (
	"encoding/json"
	"net/http"
	"practice/database"
	"practice/utils"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.IProduct

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "invalid json formate", 400)
		return
	}

	newProduct.ID = len(database.ProductList) + 1

	database.ProductList = append(database.ProductList, newProduct)

	utils.SendData(w, newProduct, 201)

}
