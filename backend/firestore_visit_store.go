package main

import (
	"context"
	"errors"
	"reflect"

	"cloud.google.com/go/firestore"
)

var (
	ErrInvalidCountValue = errors.New("invalid format for count value")
	ErrMissingCountKey = errors.New("missing count key in document data")
) 

type SnapShot interface {
	Data()  map[string]interface{}	
}

type Document interface {
	Get(context.Context) (SnapShot,  error) 
	Create(context.Context, interface{}) (*firestore.WriteResult, error)
}

type Client interface {
	Doc(path string) Document
}

type FirestoreVisitStore struct {
	client Client
}

func (f *FirestoreVisitStore) GetVisits(ctx context.Context) (int64, error) {
	ny := f.client.Doc("cloud-resume-challenge/visits")
	docsnap, err := ny.Get(ctx)
	if err != nil {
		return 0, err
	}
	
	dataMap := docsnap.Data()
	count := dataMap["count"]
	if count == nil {
		return 0, ErrMissingCountKey
	}

	visitCountValue := reflect.ValueOf(count)
	if visitCountValue.Kind() != reflect.Int64 {
		return 0, ErrInvalidCountValue
	}
	
	return visitCountValue.Int(), nil
}

