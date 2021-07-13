package main

type Config struct {
	AccessToken string `json:"access_token"`
}

type Job struct {
	Status  chan string
	Notify  chan string
	KillSig chan bool
}
