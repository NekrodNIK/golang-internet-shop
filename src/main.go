package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type Response struct {
	Title string
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("src/templates/base.html")
	if err != nil {
		panic(err)
		return
	}
	data := Response{
		Title: "MainPage",
	}
	t.Execute(w, data)
}

func main() {
	r := mux.NewRouter()
	g := r.Methods("GET").Subrouter()

	g.HandleFunc("/", mainPage)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
