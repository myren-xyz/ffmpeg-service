package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	fileName    string
	fullUrlFile string
)

func download(url string, jobID string) {
	fullUrlFile = url
	fileName = "inp.mp3"

	// Create blank file
	err := createDir(jobID)
	if err != nil {
		j := jobs[jobID]
		passToChannel(&j, "failed fetching")
		killSig(&j)
		return
	}

	file := createFile(jobID)

	// Put content on file
	err = putFile(file, httpClient())
	if err != nil {
		j := jobs[jobID]
		passToChannel(&j, "failed fetching")
		killSig(&j)
		return
	}

	j := jobs[jobID]
	passToChannel(&j, "fetched")
}

func putFile(file *os.File, client *http.Client) error {
	resp, err := client.Get(fullUrlFile)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)

	defer file.Close()

	if err != nil {
		return err
	}

	return nil
}

func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}

func createFile(path string) *os.File {
	newPath := fmt.Sprintf("./%s/", path)
	file, err := os.Create(newPath + fileName)
	checkError(err)
	return file
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
