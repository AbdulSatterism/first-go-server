package rest

import (
	"fmt"
	"net/http"
	"practice/config"
	"practice/rest/middleware"
	"strconv"
)

func Start(cnf config.Config) {
	manager := middleware.NewManager()

	manager.Use(middleware.Preflight, middleware.Cors, middleware.Logger)

	mux := http.NewServeMux()

	wrappedMux := manager.WrapMux(mux)

	InitRoutes(mux, manager)

	fmt.Println("server running on port :", cnf.Port)

	add := ":" + strconv.Itoa(cnf.Port)

	err := http.ListenAndServe(add, wrappedMux)

	if err != nil {
		fmt.Println("something went wrong server is not running", err)
	}

}
