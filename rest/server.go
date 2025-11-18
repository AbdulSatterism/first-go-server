package rest

import (
	"fmt"
	"net/http"
	"practice/config"
	"practice/rest/handlers/product"
	"practice/rest/middleware"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
}

func NewServer(cnf *config.Config, productHandler *product.Handler) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
	}
}

func (s *Server) Start() {
	manager := middleware.NewManager()

	manager.Use(middleware.Preflight, middleware.Cors, middleware.Logger)

	mux := http.NewServeMux()

	wrappedMux := manager.WrapMux(mux)

	// register routes
	s.productHandler.RegisterRoutes(mux, manager)

	fmt.Println("server running on port :", s.cnf.Port)

	add := ":" + strconv.Itoa(s.cnf.Port)

	err := http.ListenAndServe(add, wrappedMux)

	if err != nil {
		fmt.Println("something went wrong server is not running", err)
	}

}
