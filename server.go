package main

import (
	APIartist "Groupie-Tracker/src/APIartist"
	APIcarte "Groupie-Tracker/src/APIlocations"
	home "Groupie-Tracker/src/home"
	"net/http"
)

func main() {

	http.HandleFunc("/", home.Home)
	http.HandleFunc("/groupes", APIartist.ArtistPage)
	http.HandleFunc("/carte", APIcarte.CartePage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
