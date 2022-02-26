package main

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrMissingCountKey = errors.New("missing count key in document data")
)

type BadValueKindError struct {
	expected reflect.Kind
	got reflect.Kind
}

func (b BadValueKindError) Error() string {
	return fmt.Sprintf("expected kind %v, but got %v", b.expected, b.got)
}

type FirestoreVisitStore struct {
	client Client
}

func (fvs *FirestoreVisitStore) GetVisits(ctx context.Context) (int64, error) {
	visits := fvs.client.Doc("cloud-resume-challenge/visits")
	docsnap, err := visits.Get(ctx)
	
	if isNotFoundError(err) {
		return 0, nil
	}

	if err != nil {

		return 0, err
	}
	
	dataMap := docsnap.Data()
	count := dataMap["count"]
	if count == nil {
		return 0, ErrMissingCountKey
	}
	
	countValue := reflect.ValueOf(count)
	if countValue.Kind() != reflect.Int64 {
		return 0, BadValueKindError{reflect.Int64, countValue.Kind()}
	}
	
	return countValue.Int(), nil
}

func(fvs *FirestoreVisitStore) RecordVisit(ctx context.Context) error {
	visits := fvs.client.Doc("cloud-resume-challenge/visits")

	currentCount, err := fvs.GetVisits(ctx)
	if err != nil {
		return err
	}
	
	_, err = visits.Set(ctx, map[string]int64{
		"count": currentCount + 1,
	})
	if err != nil {
		return err
	}

	return nil
}

func isNotFoundError(err error) bool {
	if s, ok := status.FromError(err); ok {
		return s.Code() == codes.NotFound
	} 

	return false
}