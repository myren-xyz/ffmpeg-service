package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var config Config
var jobs map[string]Job

func init() {
	jobs = make(map[string]Job)
	file, err := ioutil.ReadFile("./.config.json")
	if err != nil {
		fmt.Printf("error in reading json config file: %s\n", err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Printf("error in unamarshalling: %s", err)
	}
}

func main() {
	prune()
	http.HandleFunc("/api/v1/convert", convertRoute)
	http.HandleFunc("/api/v1/subscribe", subscribe)
	http.ListenAndServe(":8080", nil)
}
