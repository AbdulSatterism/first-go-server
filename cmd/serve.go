package cmd

import (
	"fmt"
	"net/http"

	"practice/middleware"
)

func Serve() {
	manager := middleware.NewManager()

	manager.Use(middleware.Preflight, middleware.Cors, middleware.Logger)

	mux := http.NewServeMux()

	wrappedMux := manager.WrapMux(mux)

	InitRoutes(mux, manager)

	fmt.Println("server running on port :3000")

	err := http.ListenAndServe(":3000", wrappedMux)

	if err != nil {
		fmt.Println("something went wrong server is not running", err)
	}

}
