package LoadAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

const JSONlocation string = "https://groupietrackers.herokuapp.com/api/locations/"

type Location struct {
	Index     string   `json:"id"`
	Locations []string `json:"locations"`
}

type PageDataLocations struct {
	Locations Location
}

func CartePage(w http.ResponseWriter, r *http.Request) {

	locations := GetAPI()
	pageCarte := PageDataLocations{Locations: locations}
	ts, err := template.ParseFiles("./HTML/layout.html", "./HTML/carte.html", "./HTML/navbar.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error - PARSING", 500)
		return
	}
	err = ts.Execute(w, pageCarte)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error - EXECUTE", 500)
	}
}

func GetAPI() Location {
	var PlaceList Location
	for i := 0; i <= 52; i++ {

		res, err := http.Get(JSONlocation + strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
		}
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal(data, &PlaceList)
	}
	return PlaceList
}
