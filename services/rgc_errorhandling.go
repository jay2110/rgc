package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type APIError interface {
	APIError() (int, string)
}

type APIErrorStruct struct {
	status int
	err    error
}

func (e APIErrorStruct) APIError() (int, error) {
	return e.status, e.err
}

var (
	ErrAuth       = &APIErrorStruct{status: http.StatusUnauthorized, err: errors.New("Authorizatin failed")}
	ErrServerDown = &APIErrorStruct{status: http.StatusServiceUnavailable, err: errors.New("Upstream server is temporarly unavailable")}
	ErrValidation = APIErrorStruct{status: http.StatusBadRequest, err: errors.New("co-ordinates should be in numeric values within range Lat(-90,90), long(-180,180) ")}
)

func JSONHandleError(w http.ResponseWriter, apiErr APIErrorStruct) {
	switch {
	case apiErr.status == 0 && apiErr.err != nil:
		w.WriteHeader(ErrValidation.status)
		json.NewEncoder(w).Encode(string(ErrValidation.err.Error()))
	case apiErr.status == 401:
		w.WriteHeader(ErrAuth.status)
		json.NewEncoder(w).Encode(string(ErrAuth.err.Error()))
	case apiErr.status > 401:
		w.WriteHeader(ErrServerDown.status)
		json.NewEncoder(w).Encode(string(ErrServerDown.err.Error()))
	}

}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
