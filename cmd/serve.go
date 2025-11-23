package cmd

import (
	"fmt"
	"os"
	"practice/config"
	"practice/db"
	"practice/repo"
	"practice/rest"
	"practice/rest/handlers/product"
	"practice/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo(dbCon)

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)

	server := rest.NewServer(cnf, productHandler)

	server.Start()
}
