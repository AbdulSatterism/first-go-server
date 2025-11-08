package main

import (
	"fmt"
	"net/http"
	"strings"
)

func introduceMe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Myself Abdul Satter")
}

func aboutMe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "about me")
}
func dynamicRoute(w http.ResponseWriter, r *http.Request) {

	id := strings.Split(r.URL.Path, "/")[2]

	fmt.Fprintln(w, "dynamic route: ", id)
}

func main() {

	mux := http.NewServeMux() // this is route

	mux.HandleFunc("/about", introduceMe)

	mux.HandleFunc("/home", aboutMe)
	mux.HandleFunc("/home/{id}", dynamicRoute)

	fmt.Println("server running on port :3000")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("something went wrong server is not runnign", err)
	}

}
