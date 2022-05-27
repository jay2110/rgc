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

const FileName = "../env/configuration.yaml"

type GeoCode interface {
	ReverseGeoCodeImpl(models.InputData) (models.Output, APIErrorStruct)
	BaseUrl(models.InputData) string
}

// ReverseGeoCoder godoc
// @Summary requesting for address
// @Description Create a new adress with the input paylod
// @Tags
// @Accept  json
// @Produce  json
// @Param input body Input true "Create address request"
// @Success 200 {object} Output
// @Router /position [post]
func ReverseGeoCoder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var input models.InputData
	var apiErr APIErrorStruct
	json.NewDecoder(r.Body)
	apiErr.err = decoder.Decode(&input)
	if apiErr.err != nil {
		JSONHandleError(w, apiErr)
	} else {
		data, apiErr := ReverseGeoCodeImpl(input)
		JSONHandleError(w, apiErr)
		json.NewEncoder(w).Encode(&data)
	}
}

func BaseUrl(input models.InputData) string {
	conf := Connection(FileName)
	err := conf.Validate()
	CheckError(err)
	url := conf.Url + conf.Apikey + "&at=" + fmt.Sprint(input.Latitude) + "," + fmt.Sprint(input.Longitude)
	return url
}

func Connection(FileName string) config.ConfigStruct {
	confContent, err := ioutil.ReadFile(FileName)
	CheckError(err)
	conf := config.ConfigStruct{}
	if err := yaml.Unmarshal(confContent, &conf); err != nil {
		log.Println(err)
	}
	return conf
}

func ReverseGeoCodeImpl(input models.InputData) (models.Output, APIErrorStruct) {
	var data models.Output
	var res *http.Response
	var apiError APIErrorStruct
	apiError.err = input.Validate()
	if apiError.err != nil {
		return data, apiError
	} else {
		res, apiError.err = http.Get(BaseUrl(input))
		apiError.status = res.StatusCode
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &data)
		return data, apiError
	}

}
