package main

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestGetVisitsFromFirestore( t *testing.T) {
	cases := []int64{1, 2, 3}

	for _, visitCount := range cases {
		t.Run(fmt.Sprintf("get visit count %d from client", visitCount), func(t *testing.T) {
			client := makeMockFirestoreClient("count", visitCount)
			ctx := context.Background()
			store := FirestoreVisitStore{client}
	
			got, err := store.GetVisits(ctx)
	
			if err != nil {
				t.Fatalf("expected no error but got %v", err)
			}

			if got != visitCount {
				t.Errorf("got %d want %d", got, visitCount)
			}
		})
	}

	t.Run("returns error if snapshot cannot be retrieved", func(t *testing.T) {
		client := makeMockFirestoreClient("count", 0)
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			store := FirestoreVisitStore{client}
	
			_, err := store.GetVisits(ctx)
		
			if err == nil {
				t.Fatal("expected an error")
			}
	})

	t.Run("returns error if count value is not valid int", func(t *testing.T) {
		client := makeMockFirestoreClient("count", "invalid")
			ctx := context.Background()
			store := FirestoreVisitStore{client}
	
			_, err := store.GetVisits(ctx)
		
			if err != ErrInvalidCountValue {
				t.Errorf("expected error %v got %v", ErrInvalidCountValue, err)
			}
	})

	t.Run("returns error if count key is not found", func(t *testing.T) {
		client := makeMockFirestoreClient("notFoundKey", 10)
		ctx := context.Background()
		store := FirestoreVisitStore{client}

		_, err := store.GetVisits(ctx)
	
		if err != ErrMissingCountKey {
			t.Errorf("expected error '%v' got '%v'", ErrMissingCountKey, err)
		}
	})
}

// func TestRecordVisitOnFirestore( t *testing.T) {
// 	t.Run("record new visit in firestore", func(t *testing.T) {
// 		client := makeMockFirestoreClient("count", 0)
// 		ctx := context.Background()
// 		store := FirestoreVisitStore{client}

// 		store.RecordVisit(ctx)
	
// 		if client.GetRecordedVisitCount() != 1 {
// 			t.Errorf("got %d want %d", client.GetRecordedVisitCount(), 1)
// 		}
// 	})
// }

type StubSnapShot struct {
	data interface{}
}

func (ms *StubSnapShot) Data() map[string]interface{} {
	mappedData := make(map[string]interface{})
	
	v := reflect.ValueOf(ms.data)
	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			mappedData[key.String()] = v.MapIndex(key)
		}
	}	

	return mappedData
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
	mf.snp = &StubSnapShot{data}
	return nil, nil
}

type MockFirestoreClient struct {
	doc Document
}

func (m *MockFirestoreClient) Doc(path string) Document {
	return m.doc 
}

func (m *MockFirestoreClient) GetRecordedVisitCount() int64 {
	snapShot, _ := m.doc.Get(context.Background())
	data := snapShot.Data()
	return reflect.ValueOf(data["count"]).Int()
}

func makeMockFirestoreClient (dataKey string, visitCount interface{}) *MockFirestoreClient {
	snapShotMap := make(map[string]interface{})
	snapShotMap[dataKey] = visitCount
	
	doc := &StubFirestoreDoc{
		&StubSnapShot{
			snapShotMap,
		},
	}

	return &MockFirestoreClient{doc}
}



	// v := reflect.ValueOf(data)
	// if v.Kind() != reflect.Map {
	// 	return nil, fmt.Errorf("err creating snapshot")
	// }
		
	// var newMap map[string]interface{}
	// for _, key := range v.MapKeys() {
	// 	newMap[key.String()] = v.MapIndex(key)
	// }
	
	// mf.snp = &SpySnapShot{newMap}