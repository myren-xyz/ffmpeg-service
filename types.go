package main

import (
	"encoding/json"
)

type Config struct {
	AccessToken string `json:"access_token"`
}

type Job struct {
	Status  chan string
	Notify  chan string
	KillSig chan bool
}

type Response struct {
	OK    bool   `json:"ok"`
	JobID string `json:"job_id"`
}

func (r *Response) toStr() string {
	json, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(json)
}
