package Home

import (
	"log"
	"net/http"
	"text/template"
)

var t *template.Template
var tErr *template.Template

func Home(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("HTML/layout.html", "HTML/home.html", "HTML/navbar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, nil)
}
