package main

import (
	APIartist "Groupie-Tracker/src/APIartist"
	APIcarte "Groupie-Tracker/src/APIlocations"
	accueil "Groupie-Tracker/src/accueil"
	"net/http"
)

func main() {

	http.HandleFunc("/", accueil.AccueilPage)
	http.HandleFunc("/groupes", APIartist.ArtistPage)
	http.HandleFunc("/carte", APIcarte.CartePage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
