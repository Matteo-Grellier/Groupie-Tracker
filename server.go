package main

import (
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title    string
	Articles []string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// Création d'une page
	p := Page{"Titre de ma page", []string{"Item 1", "Item 2", "Item 3"}}

	// Création d'une nouvelle instance de template
	t := template.New("Label de ma template")

	// Déclaration des fichiers à parser
	t = template.Must(t.ParseFiles("HTML/index.html", "HTML/groupes.html"))

	// Exécution de la fusion et injection dans le flux de sortie
	// La variable p sera réprésentée par le "." dans le layout
	// Exemple {{.}} == p
	err := t.ExecuteTemplate(w, "indexPage", p)

	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
