package main

import (
	"fmt"
	"log"
	"net/http"
)

//a server created in golang

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//check url
	if r.URL.Path != "/hello" {
		http.Error(w, "error 404 not found", http.StatusNotFound)
		return
	}

	//check path
	if r.Method != "GET" {
		http.Error(w, "method is not suported", http.StatusNotFound)
		return
	}

	//write to response
	fmt.Fprintf(w, "hello!")

}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting a server on port 4000")

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
