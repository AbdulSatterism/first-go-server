package product

import (
	"net/http"
	"practice/database"
	"practice/utils"
	"strconv"
)

func (h *Handler) GetProductsById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pId, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "please provide valid product id", 400)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == pId {
			utils.SendData(w, product, 200)
		}
	}

	utils.SendData(w, "product not found ", 404)
}
