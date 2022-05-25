package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jay2110/rgc.git/services"
)

func HandleRequest() {
	http.HandleFunc("/position", services.ReverseGeoCoder)
	conf := services.Connection(services.FileName)
	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(conf.Server.Port), nil))
}
