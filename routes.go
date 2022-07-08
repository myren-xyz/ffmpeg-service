package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func convertRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileUrl := vars["file-path"]
	uploadPath := vars["upload-path"]
	issuer := vars["issuer"]

	fileNameSlice := strings.Split(fileUrl, ".")
	fileExt := fileNameSlice[len(fileNameSlice)-1]

	// tmpMAID cookie
	tmpMAIDcookie, err := r.Cookie("tmpMAID")
	if err != nil {
		// return unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// issuer should be passed when converting and uploading has been finished
	// should return job id
	newJobID := generateSeq(6)
	job := Job{
		Status:  make(chan string, 1),
		Notify:  make(chan string, 1),
		KillSig: make(chan bool, 1),
	}

	go startAct(fileUrl, newJobID, issuer, uploadPath, tmpMAIDcookie, fileExt)

	jobs.Lock()
	jobs.store[newJobID] = job
	jobs.Unlock()

	jobs.RLock()
	j := jobs.store[newJobID]
	jobs.RUnlock()

	passToChannel(&j, "inq")

	res := Response{
		OK:    true,
		JobID: newJobID,
	}

	// return http created status
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, res.toStr())
}

func subscribe(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	vars := mux.Vars(r)

	jobID := vars["job-id"]
	// if jobs[jobID] == (Job{}) {
	// 	// no jobs availavle with this jobID
	// }
	select {
	case e := <-jobs.store[jobID].Notify:
		mar, _ := json.Marshal(e)
		fmt.Fprintf(w, "data: %s\n\n", string(mar))
	}

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}
