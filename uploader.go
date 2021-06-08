package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func upload() {
	files, err := ioutil.ReadDir("./temp")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		uploadSingle(file.Name())
	}
}

func uploadSingle(path string) {
	file, err := os.Open("./temp/" + path)
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
