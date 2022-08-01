package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	fmt.Print("Starting Server...")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")

	name := r.FormValue("name")
	age := r.FormValue("age")
	address := r.FormValue("address")
	salary := r.FormValue("salary")

	newPerson := Person{
		name:    name,
		age:     age,
		address: address,
		salary:  salary,
	}

	fmt.Fprintf(w, "%v", newPerson)
}
