package Artists

import (
	"html/template"
	"log"
	"net/http"
)

func Artists(w http.ResponseWriter, r *http.Request) {

	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("HTML/layout.html", "HTML/artist.html", "HTML/navbar.html")

	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, nil)
}
