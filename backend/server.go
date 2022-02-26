package main

import (
	"fmt"
	"log"
	"net/http"
)


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
	err := v.store.RecordVisit(r.Context())
	if err != nil {
		v.handleRouterError(w, err)
		return
	}
}

func (v *VisitCountServer) GetVisits(w http.ResponseWriter, r *http.Request)  {
	currentVisits, err := v.store.GetVisits(r.Context())
	if err != nil {
		v.handleRouterError(w, err)		
		return
	}

	fmt.Fprint(w, currentVisits)
}

func (v *VisitCountServer) handleRouterError(w http.ResponseWriter, err error) {
	log.Printf("Something went wrong\n %v\n", err)
	w.WriteHeader(http.StatusInternalServerError)
}

func NewVisitCountServer (client Client) *VisitCountServer {
	return &VisitCountServer{&FirestoreVisitStore{client}}
}
