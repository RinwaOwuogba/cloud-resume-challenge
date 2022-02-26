package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)


type StubVisitStore struct {
	visits int	
}

func (s *StubVisitStore) GetVisits() (int, error) {
	return s.visits, nil
}

func (s *StubVisitStore) RecordVisit() {
	s.visits++
}

func TestRecordVisit(t *testing.T) {
	t.Run("increase visit count by 1", func(t *testing.T) {
		store := &StubVisitStore{1}
		server := &VisitCountServer{store}

		request := newRecordVisitRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		
		assertStatus(t, response.Result().StatusCode, http.StatusOK)

		got, _ := store.GetVisits()
		assertVisitCount(t, fmt.Sprintf("%d",got ), "2")
	})

	t.Run("return 404 for invalid route", func(t *testing.T) {
		store := &StubVisitStore{}
		server := &VisitCountServer{store}

		request, _ := http.NewRequest(http.MethodPut, "/api/visitsss", nil)
		response := httptest.NewRecorder();

		server.ServeHTTP(response, request)
		
		assertStatus(t, response.Result().StatusCode, http.StatusNotFound)
	})
}

func TestGetVisitCount(t *testing.T) {
	t.Run("return current visit count", func(t *testing.T) {
		store := &StubVisitStore{133}
		server := &VisitCountServer{store}

		request := newGetVisitRequest()
		response := httptest.NewRecorder();

		server.ServeHTTP(response, request)
		
		assertStatus(t, response.Result().StatusCode, http.StatusOK)
		assertVisitCount(t, response.Body.String(), "133")
	})
}


func assertStatus (t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d wanted %d", got,want)
	}
}

func assertVisitCount (t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but wanted %q visit(s)", got,want)
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
