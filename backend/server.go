package backend

import (
	"context"
	"fmt"
	"log"
	"net/http"
)


type SnapShot interface {
	Data() map[string]interface{}
}

type Document interface {
	Get(context.Context) (SnapShot, error)
	Set(context.Context, interface{}) (interface{}, error)
}

type Client interface {
	Doc(path string) Document
}

type VisitStore interface {
	GetVisits(ctx context.Context) (int64, error)
	RecordVisit(ctx context.Context) error
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

func ServerEntry(w http.ResponseWriter, r *http.Request) {
	client := GetFirestoreClient()
	server := NewVisitCountServer(MakeFirestoreClientAdapter(client))
	server.ServeHTTP(w, r)
}

