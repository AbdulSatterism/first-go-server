package cmd

import (
	"fmt"
	"os"
	"practice/config"
	"practice/db"
	"practice/rest"
	"practice/rest/handlers/product"
	"practice/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ensure database connected
	dbCon.Ping()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares)

	server := rest.NewServer(cnf, productHandler)

	server.Start()
}
