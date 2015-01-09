package main

import (
	"io"
	"log"
	"net/http"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/form", 302)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	var fd *MyForm = &MyForm{
		Age:   18,
		Token: "345625145123451234123412342345",
	}

	form, err := FormCreate(fd)

	if err != nil {
		log.Println(err)
		io.WriteString(w, "can not create form")
	} else {
		io.WriteString(w, htmlLayout(form))
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	var fd *MyForm = &MyForm{}

	err := FormRead(fd, r)

	if err == nil {
		io.WriteString(w, "saved")
	} else {
		io.WriteString(w, "error")
	}
}
