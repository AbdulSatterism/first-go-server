package cmd

import (
	"practice/config"
	"practice/rest"
	"practice/rest/handlers/product"
	"practice/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares)

	server := rest.NewServer(cnf, productHandler)

	server.Start()
}
