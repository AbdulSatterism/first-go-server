package product

import (
	"net/http"
	"practice/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts), h.middlewares.AuthenticateJWT))

	// create product
	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProduct)))

	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(h.GetProductsById)))

}
