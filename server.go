package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template
var c *template.Template

func main() {
	t = template.Must(template.ParseFiles("./HTML/index.html"))
	c = template.Must(template.ParseFiles("./HTML/projet.html"))

	http.Handle("/stuff/", http.StripPrefix("/stuff", http.FileServer(http.Dir("CSS"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/ascii-art", getAscii)
	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		fmt.Fprintf(w, "404 NOT FOUND")
		return
	}
	t.Execute(w, "index.html")
}

func getAscii(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	text := req.FormValue("first")

	l := struct {
		Phrase string
	}{
		Phrase: text,
	}
	c.ExecuteTemplate(w, "projet.html", l)
}
