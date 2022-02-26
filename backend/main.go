package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
)

type SnapShot interface {
	Data()  map[string]interface{}	
}

type Document interface {
	Get(context.Context) (SnapShot,  error) 
	Set(context.Context, interface{}) (interface{}, error)
}

type Client interface {
	Doc(path string) Document
}

type VisitStore interface {
	GetVisits(ctx context.Context) (int64, error) 
	RecordVisit(ctx context.Context) error
}


func main() {	
	projectID := os.Getenv("GCP_PROJECT")
	
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}

	store := FirestoreVisitStore{&FirestoreClient{client}}
	server := &VisitCountServer{&store}
	log.Fatal(http.ListenAndServe(":5000", server))
}