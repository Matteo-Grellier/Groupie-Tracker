package main

import (
	artist "Groupie-Tracker/src/artist"
	carte "Groupie-Tracker/src/carte"
	home "Groupie-Tracker/src/home"
	api "Groupie-Tracker/src/restAPI"
	"net/http"
)

func main() {
	http.HandleFunc("/", home.Home)
	http.HandleFunc("/artist", artist.Artists)
	http.HandleFunc("/groupes", api.GroupePage)
	http.HandleFunc("/carte", carte.Carte)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
