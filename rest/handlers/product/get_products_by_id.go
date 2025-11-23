package product

import (
	"net/http"
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

	product, err := h.productRepo.Get(pId)
	if err != nil {
		utils.SendError(w, 404, "product not found")
		return
	}

	utils.SendData(w, product, 200)
}
