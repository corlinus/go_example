package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/create", createHandler)

	log.Println("Server listening on http://0.0.0.0:4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
