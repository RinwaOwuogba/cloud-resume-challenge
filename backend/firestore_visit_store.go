package main

import (
	"context"
	"errors"
	"reflect"
)

var (
	ErrInvalidCountValue = errors.New("invalid format for count value")
	ErrMissingCountKey = errors.New("missing count key in document data")
) 


type FirestoreVisitStore struct {
	client Client
}

func (fvs *FirestoreVisitStore) GetVisits(ctx context.Context) (int, error) {
	visits := fvs.client.Doc("cloud-resume-challenge/visits")
	docsnap, err := visits.Get(ctx)
	if err != nil {
		return 0, err
	}
	
	dataMap := docsnap.Data()
	count := dataMap["count"]
	if count == nil {
		return 0, ErrMissingCountKey
	}
	
	countValue := reflect.ValueOf(count)
	if countValue.Kind() != reflect.Int {
		return 0, ErrInvalidCountValue
	}
	
	return int(countValue.Int()), nil
}

func(fvs *FirestoreVisitStore) RecordVisit(ctx context.Context) error {
	visits := fvs.client.Doc("cloud-resume-challenge/visits")

	currentCount, err := fvs.GetVisits(ctx)
	if err != nil {
		return err
	}
	
	_, err = visits.Create(ctx, map[string]int{
		"count": currentCount + 1,
	})
	if err != nil {
		return err
	}

	return nil
}
