package main

import (
	APIgroupe "Groupie-Tracker/src/APIgroupe"
	APIcarte "Groupie-Tracker/src/APIlocations"
	home "Groupie-Tracker/src/home"
	"net/http"
)

func main() {

	http.HandleFunc("/", home.Home)
	http.HandleFunc("/groupes", APIgroupe.GroupePage)
	http.HandleFunc("/carte", APIcarte.CartePage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
