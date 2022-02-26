package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingVisitsAndGettingThem(t *testing.T) {
	server := NewVisitCountServer(makeDefaultSpyFirestoreClient())
	response := httptest.NewRecorder()

	server.ServeHTTP(httptest.NewRecorder(), newRecordVisitRequest())
	server.ServeHTTP(httptest.NewRecorder(), newRecordVisitRequest())
	server.ServeHTTP(httptest.NewRecorder(), newRecordVisitRequest())

	server.ServeHTTP(response, newGetVisitRequest())

	assertStatus(t, response.Result().StatusCode, http.StatusOK)
	if response.Body.String() != "3" {
		t.Errorf("got %q want %q", response.Body.String(), "3")
	}
}