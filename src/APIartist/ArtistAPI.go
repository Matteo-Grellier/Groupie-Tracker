package LoadAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

const JSONartist string = "https://groupietrackers.herokuapp.com/api/artists"

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	CreationDate int      `json:"creationDate"`
	Members      []string `json: "members"`
	FirstAlbum   string   `json: "firstAlbum"`
}

type PageDataArtist struct {
	Artists []Artist
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {

	Artists := GetAPI()
	pageArtists := PageDataArtist{Artists: Artists}
	ts, err := template.ParseFiles("./HTML/layout.html", "./HTML/Artists.html", "./HTML/navbar.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error - PARSING", 500)
		return
	}
	err = ts.Execute(w, pageArtists)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error - EXECUTE", 500)
	}
}

func GetAPI() []Artist {

	var ArtistList []Artist
	res, err := http.Get(JSONartist)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	json.Unmarshal(data, &ArtistList)
	return ArtistList
}
