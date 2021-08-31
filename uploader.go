package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func upload(jobID string, issuedBy string, uploadPath string, cookie *http.Cookie) {

	files, err := ioutil.ReadDir("./" + jobID)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileExt := filepath.Ext(file.Name())
		// if file extension was .ts or .m3u8 then upload it
		if fileExt == ".ts" || fileExt == ".m3u8" {
			msg := "gonna be uploaded"
			log.Println(msg)
		} else {
			continue
		}

		request, err := uploadSingle(file.Name(), issuedBy, uploadPath, cookie, jobID)
		if err != nil {
			log.Println(err)
			j := jobs[jobID]
			passToChannel(&j, "failed uploading")
			killSig(&j)
			return
		}
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			log.Println(err)
			j := jobs[jobID]
			passToChannel(&j, "failed uploading")
			killSig(&j)
			return
		} else {
			var bodyContent []byte
			resp.Body.Read(bodyContent)
			resp.Body.Close()
		}
	}

	j := jobs[jobID]
	passToChannel(&j, "uploaded")

}

func uploadSingle(path string, issuedBy string, uploadPath string, cookie *http.Cookie, jobID string) (*http.Request, error) {
	fpath := fmt.Sprintf("./%s/%s", jobID, path)
	file, err := os.Open(fpath)
	if err != nil {
		log.Println(err)
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
	part, err := writer.CreateFormFile("file", fi.Name())
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
	req.Header.Add("Origin", "https://utils.myren.xyz")
	req.Header.Add("s2rj-access-token", config.AccessToken)
	req.AddCookie(cookie)
	return req, err
}
