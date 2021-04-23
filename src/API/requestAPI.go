package LoadAPI

import (
	"io/ioutil"
	"log"
	"net/http"
)

func LoadAPI(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		log.Fatal(err)
	}

	respData, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(respData)
}
