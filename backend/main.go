package main

import (
	"context"
)

type SnapShot interface {
	Data()  map[string]interface{}	
}

type Document interface {
	Get(context.Context) (SnapShot,  error) 
	Create(context.Context, interface{}) (interface{}, error)
}

type Client interface {
	Doc(path string) Document
}

type VisitStore interface {
	GetVisits() (int, error) 
	RecordVisit()
}


func main() {	
	// projectID := os.Getenv("GCLOUD_PROJECT_ID")
	
	// ctx := context.Background()
	// client, err := firestore.NewClient(ctx, projectID)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// store := FirestoreVisitStore{&FirestoreClient{client}}
	// server := &VisitCountServer{&store}
	// log.Fatal(http.ListenAndServe(":5000", server))
}