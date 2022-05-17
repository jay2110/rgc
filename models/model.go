package models

import "errors"

type InputData struct {
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
}

type Output struct {
	Items []struct {
		Title string `json:"title"`

		Id string `json:"id"`

		ResultType string `json:"resultType"`

		Address struct {
			Label string `json:"label"`

			CountryCode string `json:"countryCode"`

			CountryName string `json:"countryName"`

			StateCode string `json:"stateCode"`

			State string `json:"state"`

			Country string `json:"country"`

			City string `json:"city"`

			District string `json:"district"`

			PostalCode string `json:"postalCode"`
		} `json:"address"`

		Distance int `json:"distance"`

		Categories []struct {
			Id string `json:"id"`

			Name string `json:"name"`

			Primary bool `json:"primary"`
		} `json:"categories"`
	} `json:"items"`
}

func (a *InputData) Validate() (err error) {
	if (a.Latitude <= -90 || a.Latitude >= 90) || (a.Longitude <= -180 || a.Longitude >= 180) {
		err = errors.New("Invalid Input!!,latitude must be in between -90 and 90 and the longitude between -180 and 180")
	}
	return err
}
