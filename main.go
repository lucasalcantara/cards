package main

// Fix gcc problem https://github.com/mattn/go-sqlite3/issues/212

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("html/index.html", "html/header.html"))
	t.Execute(w, shuffleCards())
}

func newCardHandler(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("html/new.html", "html/header.html"))
	t.Execute(w, nil)
}

func toCzechHandler(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("html/toCzech.html", "html/header.html"))
	t.Execute(w, shuffleCards())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new", newCardHandler)
	http.HandleFunc("/toCzech", toCzechHandler)
	http.HandleFunc("/new/create", createHandler)
	http.HandleFunc("/remove", removeHandler)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./html/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./html/js"))))

	fmt.Println("Running application in port 8080")
	server.ListenAndServe()
}
