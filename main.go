package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Config struct {
	AccessToken string `json:"access_token"`
}

var config Config

func init() {

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
	// download("http://dl6.shirazsong.in/dl/music/99-11/Mehdi%20Jahani%20-%20Asemoone%20Mani.mp3")
	// http.HandleFunc("/api/v1/convert", convertRoute)
	// http.ListenAndServe(":80", nil)
	upload()
}

func convertRoute(w http.ResponseWriter, r *http.Request) {
	fileUrl := r.URL.Query().Get("file_url")
	err := download(fileUrl)
	if err != nil {
		fmt.Fprintf(w, "error while downloading file")
		return
	}
	err = convert()
	if err != nil {
		fmt.Fprintf(w, "error while converting file")
		return
	}
}
