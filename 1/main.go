package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform err: %v", err)
		return
	}
	fmt.Fprintf(w, "Success")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %v\naddress = %v", name, address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "wrong method", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")

}

func main() {
	fileServer := http.FileServer(http.Dir("./source"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting port at 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
		return
	}
}
