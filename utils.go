package main

import (
	"math/rand"
)

func generateSeq(length int) string {
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

func startAct(url string, jobID string) {
	for {
		select {
		case status := <-jobs[jobID].Status:
			if status == "inq" {
				go download(url, jobID)
			} else if status == "fetched" {
				go convertFile(jobID)
			} else if status == "converted" {
				go upload(jobID)
			}
		case <-jobs[jobID].KillSig:
			break
		}
	}
}
