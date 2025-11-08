package main

import (
	"fmt"
	"net/http"
)

func introduceMe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Myself Abdul Satter")
}

func aboutMe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "about me", r)
}

func main() {

	mux := http.NewServeMux() // this is route

	mux.HandleFunc("/about", introduceMe)

	mux.HandleFunc("/home/:id", aboutMe)

	fmt.Println("server running on port :3000")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("something went wrong server is not runnign", err)
	}

}
