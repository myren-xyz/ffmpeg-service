package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func upload(jobID string) {

	files, err := ioutil.ReadDir("./temp")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name() == "inp.mp3" {
			continue
		}
		request, err := uploadSingle(file.Name())
		if err != nil {
			j := jobs[jobID]
			passToChannel(&j, "failed uploading")
			killSig(&j)
			return
		}
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		} else {
			var bodyContent []byte
			resp.Body.Read(bodyContent)
			resp.Body.Close()
			fmt.Println(bodyContent)
		}
	}

	j := jobs[jobID]
	passToChannel(&j, "uploaded")

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

	req, err := http.NewRequest("POST", "http://localhost:2121/upload", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("s2rj-access-token", config.AccessToken)
	return req, err
}
