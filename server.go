package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/artist", artist)
	http.HandleFunc("/groupes", groupes)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("HTML/layout.html", "HTML/home.html", "HTML/navbar.html")

	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, nil)
}

func groupes(w http.ResponseWriter, r *http.Request) {

	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("HTML/layout.html", "HTML/groupes.html", "HTML/navbar.html")

	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, nil)
}

func artist(w http.ResponseWriter, r *http.Request) {

	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("HTML/artist.html", "HTML/layout.html")

	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, nil)
}

func carte(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("HTML/carte.html", "HTML/layout.html")

	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, nil)
}
