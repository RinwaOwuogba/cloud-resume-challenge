package main

import (
	"fmt"
	"net/http"
)

type VisitStore interface {
	GetVisits() (int, error) 
	RecordVisit()
}

type VisitCountServer struct {
	store VisitStore
}

func (v *VisitCountServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	route := "/api/visits"

	if r.URL.Path != route {
		w.WriteHeader(http.StatusNotFound)
		return
	}


	switch r.Method {
		case http.MethodPut: v.RecordVisit(w, r)
		case http.MethodGet: v.GetVisits(w, r)
	}
}

func (v *VisitCountServer) RecordVisit(w http.ResponseWriter, r *http.Request)  {
	v.store.RecordVisit()
}

func (v *VisitCountServer) GetVisits(w http.ResponseWriter, r *http.Request)  {
	currentVisits, _ := v.store.GetVisits()
	fmt.Fprint(w, currentVisits)
}