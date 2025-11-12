package cmd

import (
	"fmt"
	"net/http"
	"practice/global_router"
	"practice/handlers"
)

func Serve() {

	mux := http.NewServeMux() // this is route

	// adding new route for products
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))

	// create product
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))

	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductsById))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("server running on port :3000")

	err := http.ListenAndServe(":3000", globalRouter)

	if err != nil {
		fmt.Println("something went wrong server is not running", err)
	}

}
