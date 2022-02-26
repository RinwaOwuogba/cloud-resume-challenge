package main

import (
	"context"
	"fmt"
	"testing"
)

type StubSnapShot struct {
	data map[string]interface{}
}

func (ms *StubSnapShot) Data() map[string]interface{} {
	return ms.data
}

type StubFirestoreDoc struct {
	snp SnapShot
}

func (mf *StubFirestoreDoc) Get(ctx context.Context) (SnapShot, error) {
	if ctx.Err() != nil {
		return nil, fmt.Errorf("invalid context")
	}

	return mf.snp, nil
}

func (mf *StubFirestoreDoc) Create(ctx context.Context, data interface {}) (interface{}, error) {
	return nil, nil
}

type StubFirestoreClient struct {
	dataKey string
	value interface{}
}

func (m *StubFirestoreClient) Doc(path string) Document {
	snapShotMap := make(map[string]interface{})
	snapShotMap[m.dataKey] = m.value

	return &StubFirestoreDoc{
		&StubSnapShot{
			snapShotMap,
		},
	}
}

func TestFirestoreVisitStore( t *testing.T) {
	cases := []int64{1, 2, 3}

	for _, visitCount := range cases {
		t.Run(fmt.Sprintf("get visit count %d from client", visitCount), func(t *testing.T) {
			client := &StubFirestoreClient{"count", visitCount}
			ctx := context.Background()
			store := FirestoreVisitStore{client, ctx}
	
			got, err := store.GetVisits()
	
			if err != nil {
				t.Fatalf("expected no error but got %v", err)
			}

			if got != visitCount {
				t.Errorf("got %d want %d", got, visitCount)
			}
		})
	}

	t.Run("returns error if snapshot cannot be retrieved", func(t *testing.T) {
			client := &StubFirestoreClient{}
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			store := FirestoreVisitStore{client, ctx}
	
			_, err := store.GetVisits()
		
			if err == nil {
				t.Fatal("expected an error")
			}
	})

	t.Run("returns error if count value is not valid int", func(t *testing.T) {
			client := &StubFirestoreClient{"count", "invalid"}
			ctx := context.Background()
			store := FirestoreVisitStore{client, ctx}
	
			_, err := store.GetVisits()
		
			if err != ErrInvalidCountValue {
				t.Errorf("expected error %v got %v", ErrInvalidCountValue, err)
			}
	})

	t.Run("returns error if count key is not found", func(t *testing.T) {
		client := &StubFirestoreClient{"notFoundKey", 10}
		ctx := context.Background()
		store := FirestoreVisitStore{client, ctx}

		_, err := store.GetVisits()
	
		if err != ErrMissingCountKey {
			t.Errorf("expected error '%v' got '%v'", ErrMissingCountKey, err)
		}
})
}


	// v := reflect.ValueOf(data)
	// if v.Kind() != reflect.Map {
	// 	return nil, fmt.Errorf("err creating snapshot")
	// }
		
	// var newMap map[string]interface{}
	// for _, key := range v.MapKeys() {
	// 	newMap[key.String()] = v.MapIndex(key)
	// }
	
	// mf.snp = &StubSnapShot{newMap}