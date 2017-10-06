package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)

		cards := make([]card, 0)

		err := decoder.Decode(&cards)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		insertCards(cards)

		fmt.Fprint(w, "Phrase created!")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func removeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		id := r.Form.Get("id")

		deleteCard(id)

		fmt.Fprint(w, "Phrase removed!")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
