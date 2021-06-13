package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"mime/multipart"
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

func uploadSingle(path string) (*http.Request, error) {
	file, err := os.Open("./temp/" + path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("track", fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return http.NewRequest("POST", "http://localhost:2121/upload", body)
}
