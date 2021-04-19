package Carte

import (
	"html/template"
	"log"
	"net/http"
)

func Carte(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("HTML/layout.html", "HTML/carte.html", "HTML/navbar.html")

	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, nil)
}
