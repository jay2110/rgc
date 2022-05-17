package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jay2110/rgc.git/config"
	"github.com/jay2110/rgc.git/models"
	"gopkg.in/yaml.v2"
)

func ReverseGeoCoder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input models.InputData
	var data models.Output
	conf := Connection()
	json.NewDecoder(r.Body).Decode(&input)
	url := conf.Url + conf.Apikey + "&at=" + fmt.Sprint(input.Latitude) + "," + fmt.Sprint(input.Longitude)
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(body, &data)
	json.NewEncoder(w).Encode(&data)
}

func Connection() config.Revgeo {
	confContent, err := ioutil.ReadFile("../../application.yaml")
	if err != nil {
		panic(err)
	}
	conf := config.Revgeo{}
	if err := yaml.Unmarshal(confContent, &conf); err != nil {
		panic(err)
	}
	return conf
}
