package models

import (
	"errors"
	"log"
	"testing"
)

func TestValidate(t *testing.T) {
	var input InputData
	input.Latitude = -90.00
	input.Longitude = 2.987
	err := input.Validate()
	expectedErr := errors.New("Invalid Input!!,latitude must be in between -90 and 90 and the longitude between -180 and 180")
	if err == expectedErr {
		log.Println("validated")
	}
}
