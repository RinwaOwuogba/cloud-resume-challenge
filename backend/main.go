package main

import (
	"log"
	"net/http"
)

type InMemoryCounter struct {
	visits int
}

func (i *InMemoryCounter) GetVisits() int {
	return i.visits
}

func (i *InMemoryCounter) RecordVisit() {
	i.visits++
}

func main() {
	server := &VisitCountServer{&InMemoryCounter{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}