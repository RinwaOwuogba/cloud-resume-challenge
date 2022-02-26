package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
)

// func TestWriteResult(wr interface{}) {}
// func HelpTestWriteResult() *firestore.WriteResult {
// 	return nil
// }

func TestCreate(func() interface{}) {}
// func TestCreate(func() interface{}) {}
func HelpTestCreate () func() *firestore.WriteResult {
	return func() *firestore.WriteResult {
		return nil
	}
}


func main() {
	TestCreate(HelpTestCreate())
	// TestWriteResult(&firestore.WriteResult{})
	
	
	projectID := os.Getenv("GCLOUD_PROJECT_ID")
	
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}

	store := FirestoreVisitStore{client}
	server := &VisitCountServer{&InMemoryCounter{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}