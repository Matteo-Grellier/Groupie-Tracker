package main

import (
	api "Groupie-Tracker/src/API"
	artist "Groupie-Tracker/src/artist"
	carte "Groupie-Tracker/src/carte"
	groupe "Groupie-Tracker/src/groupes"
	home "Groupie-Tracker/src/home"
	"net/http"
)

func main() {
	http.HandleFunc("/api", api.LoadAPI)
	http.HandleFunc("/", home.Home)
	http.HandleFunc("/artist", artist.Artists)
	http.HandleFunc("/groupes", groupe.Groupes)
	http.HandleFunc("/carte", carte.Carte)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
