package main

import (
	"context"

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

type FirestoreClient struct {
	client *firestore.Client
}

func(m *FirestoreClient) Doc (path string) Document {
	return &FirestoreDocument{m.client.Doc(path)}
}
