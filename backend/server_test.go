package main

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)


type ContextFlag struct {
	string
}

var (
	DocumentGetFailFlag = ContextFlag{"DocumentGetFail"}
	DocumentGetNotFoundFlag = ContextFlag{"DocumentGetNotFound"}
	DocumentSetFailFlag = ContextFlag{"DocumentSetFail"}
	StoreFailFlag = ContextFlag{"StoreFail"}
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestGetVisitCount(t *testing.T) {
	t.Run("get current visit count", func(t *testing.T) {
		store := &StubVisitStore{133}
		server := &VisitCountServer{store}

		request := newGetVisitRequest()
		response := httptest.NewRecorder();

		server.ServeHTTP(response, request)
		
		assertStatus(t, response.Result().StatusCode, http.StatusOK)
		
		if response.Body.String() != "133" {
			t.Errorf("got %q want %q", response.Body.String(), "133")
		}
	})

	t.Run("status internal server error when store fails", func(t *testing.T) {
		store := &StubVisitStore{}
		server := &VisitCountServer{store}

		request := newGetVisitRequest()
		request = addContextFlagToRequest(request, StoreFailFlag) 
		response := httptest.NewRecorder();

		server.ServeHTTP(response, request)
		
		assertStatus(t, response.Result().StatusCode, http.StatusInternalServerError)
	
	})
}

func TestRecordVisit(t *testing.T) {
	t.Run("increase visit count by 1", func(t *testing.T) {
		store := &StubVisitStore{1}
		server := &VisitCountServer{store}

		request := newRecordVisitRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		
		got, _ := store.GetVisits(context.Background())
		
		assertStatus(t, response.Result().StatusCode, http.StatusOK)
		assertVisitCount(t, got, 2)
	})

	t.Run("status internal server error when store fails", func(t *testing.T) {
		store := &StubVisitStore{}
		server := &VisitCountServer{store}

		request := newRecordVisitRequest()
		request = addContextFlagToRequest(request, StoreFailFlag) 
		response := httptest.NewRecorder();

		server.ServeHTTP(response, request)
		
		assertStatus(t, response.Result().StatusCode, http.StatusInternalServerError)
	
	})

	t.Run("status not found for unknown route", func(t *testing.T) {
		store := &StubVisitStore{}
		server := &VisitCountServer{store}

		request, _ := http.NewRequest(http.MethodPut, "/api/visitsss", nil)
		response := httptest.NewRecorder();

		server.ServeHTTP(response, request)
		
		assertStatus(t, response.Result().StatusCode, http.StatusNotFound)
	})
}


type StubVisitStore struct {
	visits int64
}

func (s *StubVisitStore) GetVisits(ctx context.Context) (int64, error) {
	if hasContextFlag(ctx, StoreFailFlag) {
		return 0, errors.New("error getting visits")
	}
	
	return s.visits, nil
}

func (s *StubVisitStore) RecordVisit(ctx context.Context) error {
	if hasContextFlag(ctx, StoreFailFlag) {
		return errors.New("error recording visit")
	}

	s.visits++
	return nil
}

func assertStatus (t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d wanted %d", got,want)
	}
}

func assertVisitCount (t testing.TB, got, want int64) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but wanted %d visit(s)", got,want)
	}
}

func newRecordVisitRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodPut, "/api/visits", nil)
	return req
} 

func newGetVisitRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/api/visits", nil)
	return req
} 

func addContextFlagToRequest(r *http.Request, flag ContextFlag) *http.Request {
	contextWithFlag := context.WithValue(context.Background(), flag, true)
	return r.WithContext(contextWithFlag)
}