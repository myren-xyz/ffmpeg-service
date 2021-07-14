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

func upload(jobID string, issuedBy string, uploadPath string) {

	files, err := ioutil.ReadDir("./" + jobID)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name() == "inp.mp3" {
			continue
		}
		request, err := uploadSingle(file.Name(), issuedBy, uploadPath)
		if err != nil {
			j := jobs[jobID]
			passToChannel(&j, "failed uploading")
			killSig(&j)
			return
		}
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			j := jobs[jobID]
			passToChannel(&j, "failed uploading")
			killSig(&j)
			return
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

func uploadSingle(path string, issuedBy string, uploadPath string) (*http.Request, error) {
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
	url := fmt.Sprintf("https://s2rj.myren.xyz/api/v1/upload?issued_by=%s&path=%s", issuedBy, uploadPath)
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("s2rj-access-token", config.AccessToken)
	return req, err
}
