package rest

import (
	"net/http"
	"practice/rest/handlers"
	"practice/rest/middleware"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))

	// create product
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct)))

	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductsById)))

}
