package main

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestGetVisitsFromFirestore( t *testing.T) {
	cases := []int{1, 2, 3}

	for _, visitCount := range cases {
		t.Run(fmt.Sprintf("get visit count %d from client", visitCount), func(t *testing.T) {
			client := makeSpyFirestoreClient("count", visitCount)
			ctx := context.Background()
			store := FirestoreVisitStore{client}
	
			got, err := store.GetVisits(ctx)
			
			assertNoError(t, err)
			assertVisitCount(t, got, visitCount)
		})
	}

	t.Run("should return error when document get fails", func(t *testing.T) {
		client := makeDefaultSpyFirestoreClient()
		store := FirestoreVisitStore{client}
	
		_, err := store.GetVisits(getContextWithFlag(DocumentGetFailFlag))
		
		assertError(t, err)
	})

	t.Run("should return error when context is cancelled", func(t *testing.T) {
		client := makeDefaultSpyFirestoreClient()
		store := FirestoreVisitStore{client}
	
		_, err := store.GetVisits(cancelledContext())
		
		assertError(t, err)
	})


	t.Run("fail when count value is not valid int", func(t *testing.T) {
		client := makeSpyFirestoreClient("count", "invalid")
		ctx := context.Background()
		store := FirestoreVisitStore{client}
	
		_, err := store.GetVisits(ctx)
		
		assertErrorType(t, err, ErrInvalidCountValue)
	})

	t.Run("fail when count key is not found", func(t *testing.T) {
		client := makeSpyFirestoreClient("notFoundKey", 10)
		ctx := context.Background()
		store := FirestoreVisitStore{client}

		_, err := store.GetVisits(ctx)

		assertErrorType(t, err, ErrMissingCountKey)
	})
}

func TestRecordVisitOnFirestore( t *testing.T) {
	t.Run("record new visit in firestore", func(t *testing.T) {
		client := makeSpyFirestoreClient("count", 23)
		ctx := context.Background()
		store := FirestoreVisitStore{client}

		err := store.RecordVisit(ctx)
	
		assertNoError(t, err)
		assertVisitCount(t, client.GetRecordedVisitCount(), 24)
	})

	t.Run("doesn't record visit for cancelled context", func(t *testing.T) {
		client := makeDefaultSpyFirestoreClient()
		store := FirestoreVisitStore{client}

		err := store.RecordVisit(cancelledContext())
	
		assertError(t, err)
		assertVisitCount(t, client.GetRecordedVisitCount(), 0)
	})

	t.Run("doesn't record visit on client error", func(t *testing.T) {
		client := makeDefaultSpyFirestoreClient()
		ctx := getContextWithFlag(DocumentCreateFailFlag)
		store := FirestoreVisitStore{client}

		err := store.RecordVisit(ctx)

		assertError(t, err)
		assertVisitCount(t, client.GetRecordedVisitCount(), 0)
	})
}

type StubSnapShot struct {
	data map[string]interface{}
}

func (ms *StubSnapShot) Data() map[string]interface{} {
	return ms.data
}

type SpyFirestoreDoc struct {
	snp SnapShot
}

func (mf *SpyFirestoreDoc) Get(ctx context.Context) (SnapShot, error) {
	if ctx.Err() != nil {
		return nil, fmt.Errorf("invalid context")
	}

	if hasContextFlag(ctx, DocumentGetFailFlag) {
		return nil, errors.New("something made client fail")
	}

	return mf.snp, nil
}

func (mf *SpyFirestoreDoc) Create(ctx context.Context, data interface {}) (interface{}, error) {
	if hasContextFlag(ctx, DocumentCreateFailFlag) {
		return nil, errors.New("something made client fail")
	}

	mappedData := make(map[string]interface{})

	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			mappedData[key.String()] = v.MapIndex(key).Interface()
		}
	}
	mf.snp = &StubSnapShot{mappedData}
	
	return nil, nil
}

type SpyFirestoreClient struct {
	doc Document
}

func (m *SpyFirestoreClient) Doc(path string) Document {
	return m.doc 
}

func (m *SpyFirestoreClient) GetRecordedVisitCount() int {
	snapShot, _ := m.doc.Get(context.Background())
	data := snapShot.Data()
	v := reflect.ValueOf(data["count"])

	return int(v.Int())
}

func makeSpyFirestoreClient (dataKey string, visitCount interface{}) *SpyFirestoreClient {
	snapShotMap := make(map[string]interface{})
	snapShotMap[dataKey] = visitCount
	
	doc := &SpyFirestoreDoc{
		&StubSnapShot{
			snapShotMap,
		},
	}

	return &SpyFirestoreClient{doc}
}

func makeDefaultSpyFirestoreClient () *SpyFirestoreClient {
	return makeSpyFirestoreClient("count", 0)
}

func assertErrorType(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("expected error '%v' got '%v'", got, want)
	}
}

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
			t.Fatal("expected an error but got none")
		}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected no error but got %v", err)
	}
}

func hasContextFlag(ctx context.Context, flag ContextFlag) bool {
	return ctx.Value(flag) != nil 
}

func getContextWithFlag(flag ContextFlag) context.Context {
	return context.WithValue(context.Background(), flag, true)		
}

func cancelledContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}