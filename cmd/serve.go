package cmd

import (
	"fmt"
	"net/http"
	"practice/global_router"
	"practice/middleware"
)

func Serve() {
	manager := middleware.NewManager()

	manager.Use(middleware.Logger)

	mux := http.NewServeMux() // this is route

	InitRoutes(mux, manager)

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("server running on port :3000")

	err := http.ListenAndServe(":3000", globalRouter)

	if err != nil {
		fmt.Println("something went wrong server is not running", err)
	}

}
