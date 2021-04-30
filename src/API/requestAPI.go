package LoadAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

const JSON string = "https://groupietrackers.herokuapp.com/api/artists"

type Groupe struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
}

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
	res, err := http.Get(JSON)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	json.Unmarshal(data, &GroupeList)
	return GroupeList
}
