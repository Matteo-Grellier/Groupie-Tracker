package LoadAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

const JSONlocation string = "https://groupietrackers.herokuapp.com/api/locations/1"

type Location struct {
	Index     string     `json:"id"`
	Locations []Location `json:"locations"`
}

type PageDataLocations struct {
	Locations []string
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

func GetAPI() []string {

	var PlaceList []string
	res, err := http.Get(JSONlocation)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	json.Unmarshal([]byte(data), &PlaceList)
	fmt.Println(PlaceList)
	return PlaceList

}
