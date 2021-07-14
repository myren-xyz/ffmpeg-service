package main

import (
	"fmt"
	"math/rand"
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

func startAct(url string, jobID string, issuedBy string, uploadPath string) {
	for {
		select {
		case status := <-jobs[jobID].Status:
			if status == "inq" {
				go download(url, jobID)
			} else if status == "fetched" {
				go convertFile(jobID)
			} else if status == "converted" {
				go upload(jobID, issuedBy, uploadPath)
			}
		case <-jobs[jobID].KillSig:
			prune()
			break
		}
	}
}

func prune() {
	cmd := exec.Command("/bin/sh", "-c", "rm -rf ./temp/*")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}
