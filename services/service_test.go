package services

import (
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jay2110/rgc.git/models"
	"github.com/stretchr/testify/assert"
)

type ReverseGeoCodeImplTest struct {
	input  models.InputData
	apiErr APIErrorStruct
}

var (
	Input1  = models.InputData{Latitude: 47.8584, Longitude: 2.2945}
	Input2  = models.InputData{Latitude: -90.8584, Longitude: 2.2945}
	Input3  = models.InputData{Latitude: 100.8584, Longitude: 2.2945}
	Input4  = models.InputData{Latitude: 48.8584, Longitude: -180.2945}
	Input5  = models.InputData{Latitude: 48.8584, Longitude: 181.2945}
	apiErr1 = APIErrorStruct{status: 200, err: nil}
	apiErr2 = APIErrorStruct{status: 0, err: errors.New("Invalid Input!! latitude must be in between -90 and 90 and the longitude between -180 and 180")}
)

var ReverseGeoCodeImplTests = []ReverseGeoCodeImplTest{
	{Input1, apiErr1},
	{Input2, apiErr2},
	{Input3, apiErr2},
	{Input4, apiErr2},
	{Input5, apiErr2},
}

func TestReverseGeoCodeImpl(t *testing.T) {

	for _, test := range ReverseGeoCodeImplTests {
		if _, output := ReverseGeoCodeImpl(test.input); output == test.apiErr {
			t.Log("Test pass")
		}
	}
}

func TestReverseGeoCoder(t *testing.T) {
	req := httptest.NewRequest(
		"POST",
		"/position",
		strings.NewReader(`{ "Latitude":48.8584, "Longitude":2.2945}`))
	rr := httptest.NewRecorder()
	ReverseGeoCoder(rr, req)
	res := rr.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
func TestBaseurl(t *testing.T) {

	var input models.InputData
	url := BaseUrl(input)
	if len(url) == 0 {
		log.Fatal("Empty url")
	}

}

func TestConnection(t *testing.T) {
	conf := Connection("../env/configuration.yaml")
	if len(conf.Apikey) == 0 {
		log.Fatal("No apikey present")
	}
	if len(conf.Url) == 0 {
		log.Fatal("No url available")
	}
	if conf.Server.Port == 0 {
		log.Fatal("port number not present")

	}
}
