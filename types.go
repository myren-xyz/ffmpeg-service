package main

type Config struct {
	AccessToken string `json:"access_token"`
}

type Job struct {
	ID     uint
	Status string
	Events chan *Event
}

type Event struct {
	Message string
}
