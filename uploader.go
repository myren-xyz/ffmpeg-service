package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func upload() {
	file, err := os.Open("./filename.ext")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res, err := http.Post("http://127.0.0.1:5050/upload", "multipart/form-data", file)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	message, _ := ioutil.ReadAll(res.Body)
	fmt.Printf(string(message))
}
