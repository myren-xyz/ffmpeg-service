package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func convertRoute(w http.ResponseWriter, r *http.Request) {
	fileUrl := r.URL.Query().Get("file_url")
	// issuer should be passed when converting and uploading has been finished
	// should return job id
	_ = r.URL.Query().Get("issuer")
	newJobID := generateSeq(6)
	job := Job{
		Status:  make(chan string, 1),
		Notify:  make(chan string, 1),
		KillSig: make(chan bool, 1),
	}

	go startAct(fileUrl, newJobID)

	jobs[newJobID] = job
	j := jobs[newJobID]
	passToChannel(&j, "inq")
	fmt.Fprintf(w, newJobID)

}

func subscribe(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	jobID := r.URL.Query().Get("job_id")
	// if jobs[jobID] == (Job{}) {
	// 	// no jobs availavle with this jobID

	// }
	select {
	case e := <-jobs[jobID].Notify:
		mar, _ := json.Marshal(e)
		fmt.Fprintf(w, "data: %s\n\n", string(mar))
		fmt.Printf("%v\n", string(mar))
	}

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}
