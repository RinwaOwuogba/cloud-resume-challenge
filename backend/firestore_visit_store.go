package main

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrInvalidCountValue = errors.New("invalid format for count value")
	ErrMissingCountKey = errors.New("missing count key in document data")
) 

type Visit struct {
	count int64
}

type FirestoreVisitStore struct {
	client Client
}

func (fvs *FirestoreVisitStore) GetVisits(ctx context.Context) (int64, error) {
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
	cv, ok := count.(int64)
	
	
	fmt.Printf("log.Logger: %v\n", count)
	fmt.Printf("log.Logger: %v\n", countValue.Kind())
	fmt.Printf("log.Logger: %v\n", cv)
	fmt.Printf("log.Logger: %v\n", ok)
	fmt.Printf("\n\n")
	
	if countValue.Kind() != reflect.Int64 {
		return 0, ErrInvalidCountValue
	}
	
	return countValue.Int(), nil
}

func(fvs *FirestoreVisitStore) RecordVisit(ctx context.Context) {
	visits := fvs.client.Doc("cloud-resume-challenge/visits")

	currentCount, _ := fvs.GetVisits(ctx)
	visits.Create(ctx, map[string]int64{
		"count": currentCount + 1,
	})
}
