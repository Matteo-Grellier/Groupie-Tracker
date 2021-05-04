package Home

import (
	"log"
	"net/http"
	"text/template"
)

var t *template.Template
var tErr *template.Template

func Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		error(w, r, http.StatusNotFound)
		return
	}
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("HTML/layout.html", "HTML/home.html", "HTML/navbar.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	t.Execute(w, nil)
}
func error(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		//fmt.Fprint(w, "Error 404 : Page not found")
		tErr.Execute(w, template.HTML(`<h1>Error 404 : Page not found...<h1>`))
	}

	if status == 400 {
		tErr.Execute(w, template.HTML(`<h1>Error 400 : Bad Request...<h1>`))
	}

	if status == 500 {
		tErr.Execute(w, template.HTML(`<h1>Error 500 : Internal Server Error...<h1>`))
	}
}
