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

type Groupe struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	CreationDate int      `json:"creationDate"`
	Members      []string `json: "members"`
	FirstAlbum   string   `json: "firstAlbum"`
}

// var bucket *gocb.Bucket

// func SearchEndpoint(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")
// 	params := r.URL.Query()
// 	query := gocb.NewSearchQuery("groupes", cbft.NewMatchQuery(params.Get("query")))
// 	query.Fields("artist")
// 	result, _ := bucket.ExecuteSearchQuery(query)
// 	var artist []music
// 	for _, artist := range result.Hits() {
// 		artist = append(artist, Song{
// 			Id:     hit.Id,
// 			Score:  hit.Score,
// 			Artist: hit.Fields["artist"],
// 			Title:  hit.Fields["title"],
// 		})
// 	}
// 	json.NewEncoder(w).Encode(hits)
// }

type PageDataGroupe struct {
	Groupes []Groupe
}

func GroupePage(w http.ResponseWriter, r *http.Request) {

	groupes := GetAPI()
	pageGroupes := PageDataGroupe{Groupes: groupes}
	ts, err := template.ParseFiles("./HTML/layout.html", "./HTML/groupes.html", "./HTML/navbar.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error - PARSING", 500)
		return
	}
	err = ts.Execute(w, pageGroupes)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error - EXECUTE", 500)
	}
}

func GetAPI() []Groupe {

	var GroupeList []Groupe
	res, err := http.Get(JSONartist)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	json.Unmarshal(data, &GroupeList)
	return GroupeList
}
