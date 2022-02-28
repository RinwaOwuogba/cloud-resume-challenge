package backend

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

type FirestoreDocument struct {
	doc *firestore.DocumentRef
}

func (mfd *FirestoreDocument) Get(ctx context.Context) (SnapShot,  error)  {
	return mfd.doc.Get(ctx)
}

func (mfd *FirestoreDocument) Set(ctx context.Context, data interface{}) (interface{}, error) {
	return mfd.doc.Set(ctx, data)
}

type FirestoreClientAdapter struct {
	client *firestore.Client
}

func(m *FirestoreClientAdapter) Doc (path string) Document {
	return &FirestoreDocument{m.client.Doc(path)}
}

func MakeFirestoreClientAdapter (client *firestore.Client) *FirestoreClientAdapter {
	return &FirestoreClientAdapter{client}
}

func GetFirestoreClient() *firestore.Client {
	projectID := os.Getenv("GCP_PROJECT")

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

