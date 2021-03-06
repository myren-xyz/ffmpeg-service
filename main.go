package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var config Config
var jobs = struct {
	sync.RWMutex
	store map[string]Job
}{store: make(map[string]Job)}

func init() {
	file, err := ioutil.ReadFile("./.config.json")
	if err != nil {
		log.Printf("error in reading json config file: %s\n", err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Printf("error in unamarshalling: %s", err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ffmpeg/api/v1/convert/{issuer}/{file-path}/{upload-path}", cors(convertRoute)).Methods("PUT")
	r.HandleFunc("/ffmpeg/api/v1/subscribe/{job-id}", cors(subscribe)).Methods("GET")
	http.ListenAndServe(":8080", nil)
}
