package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func introduceMe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Myself Abdul Satter")
}

func aboutMe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "about me")
}

// get all products

// product struct
type IProduct struct {
	ID          int `json:"id"` // it converted into small leter id
	Name        string
	Price       float64
	Stock       int
	Description string
}

var productList []IProduct

func getProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "This method not allow", 400)
		return
	}

	encode := json.NewEncoder(w)

	encode.Encode(productList)

}

func main() {

	mux := http.NewServeMux() // this is route

	mux.HandleFunc("/about", introduceMe)

	mux.HandleFunc("/home", aboutMe)

	// adding new route for products

	mux.HandleFunc("/products", getProducts)

	fmt.Println("server running on port :3000")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("something went wrong server is not running", err)
	}

}

func init() {

	// some demo product

	product1 := IProduct{
		ID:          1,
		Name:        "Laptop",
		Price:       45000.00,
		Stock:       5,
		Description: "This is a gaming laptop",
	}

	product2 := IProduct{
		ID:          2,
		Name:        "Mobile Phone",
		Price:       25000.00,
		Stock:       10,
		Description: "This is a smartphone",
	}

	product3 := IProduct{
		ID:          3,
		Name:        "Headphones",
		Price:       5000.00,
		Stock:       15,
		Description: "This is a wireless headphone",
	}

	productList = append(productList, product1, product2, product3)

}
