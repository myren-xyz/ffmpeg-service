package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func generateSeq(length int) string {
	rand.Seed(time.Now().UnixNano())
	seq := "1324657890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	gen := ""
	for i := 0; i < length; i++ {
		max := len(seq) - 1
		rnd := rand.Intn(max)
		gen += string(seq[rnd])
	}
	return gen
}

func passToChannel(job *Job, data string) {
	go func() {
		job.Status <- data
	}()
	go func() {
		job.Notify <- data
	}()
}

func killSig(job *Job) {
	go func() {
		job.KillSig <- true
	}()
}

func startAct(url string, jobID string, issuedBy string, uploadPath string, cookie *http.Cookie, fileExt string) {
	for {
		select {
		case status := <-jobs.store[jobID].Status:
			if status == "inq" {
				go download(url, jobID)
			} else if status == "fetched" {
				go convertFile(jobID, fileExt)
			} else if status == "converted" {
				go upload(jobID, issuedBy, uploadPath, cookie)
			} else if status == "uploaded" {
				go prune(jobID)
				jobs.Lock()
				delete(jobs.store, jobID)
				jobs.Unlock()
			}
		case <-jobs.store[jobID].KillSig:
			prune(jobID)
			jobs.Lock()
			delete(jobs.store, jobID)
			jobs.Unlock()
		}
	}
}

func prune(path string) {
	cmdStr := fmt.Sprintf("rm -rf ./%s", path)
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}

func createDir(name string) error {
	err := os.Mkdir(name, 0777)
	if err != nil {
		return err
	}
	return nil
}
