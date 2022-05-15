package handlers

import (
	"REVGEOCOD/services"
	"fmt"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/position", services.ReverseGeoCoder)
	conf := services.Connection()
	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(conf.Server.Port), nil))
}
