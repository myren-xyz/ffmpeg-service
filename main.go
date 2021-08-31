package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var config Config
var jobs map[string]Job

func init() {
	jobs = make(map[string]Job)
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
	http.HandleFunc("/ffmpeg/api/v1/convert", cors(convertRoute))
	http.HandleFunc("/ffmpeg/api/v1/subscribe", cors(subscribe))
	http.ListenAndServe(":8080", nil)
}
