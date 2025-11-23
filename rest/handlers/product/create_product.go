package product

import (
	"encoding/json"
	"net/http"
	"practice/repo"
	"practice/utils"
)

type RequestCreateProduct struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req RequestCreateProduct

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "invalid json formate", 400)
		return
	}

	createProduct, err := h.productRepo.Create(repo.Product{
		Name:        req.Name,
		Price:       req.Price,
		Stock:       req.Stock,
		Description: req.Description,
	})

	if err != nil {
		http.Error(w, "failed to create product", 500)
		return
	}

	utils.SendData(w, createProduct, http.StatusCreated)

}
