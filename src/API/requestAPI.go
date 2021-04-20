package LoadAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {
	Id           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int64    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Posts []Post

func LoadAPI() {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		log.Fatal(err)
	}

	content, _ := ioutil.ReadAll(resp.Body)

	var artists Posts

	err = json.Unmarshal(content, &artists)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(artists[0].Name)
}
