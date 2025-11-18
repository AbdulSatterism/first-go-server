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

	product := database.Get(pId)

	if product == nil {
		utils.SendError(w, 404, "product not found")
		return
	}

	utils.SendData(w, product, 200)
}
