package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingVisitsAndGettingThem(t *testing.T) {
	t.Run("get visits", func(t *testing.T) {
		server := NewVisitCountServer(makeDefaultSpyFirestoreClient())
		response := httptest.NewRecorder()

		server.ServeHTTP(httptest.NewRecorder(), newRecordVisitRequest())
		server.ServeHTTP(httptest.NewRecorder(), newRecordVisitRequest())
		server.ServeHTTP(response, newRecordVisitRequest())
	
		server.ServeHTTP(response, newGetVisitRequest())
	
		assertStatus(t, response.Result().StatusCode, http.StatusOK)
		if response.Body.String() != "3" {
			t.Errorf("got %q want %q", response.Body.String(), "3")
		}
	})
	t.Run("status 500 on client get fail", func(t *testing.T) {
		server := NewVisitCountServer(makeDefaultSpyFirestoreClient())
		response := httptest.NewRecorder()
		
		request := addContextFlagToRequest(newGetVisitRequest(), DocumentGetFailFlag) 
		
		server.ServeHTTP(response, request)

		assertStatus(t, response.Result().StatusCode, http.StatusInternalServerError)
		assertResponseBodyLength(t, response.Body.Len(), 0)
	})

	t.Run("status 500 on client set fail", func(t *testing.T) {
		server := NewVisitCountServer(makeDefaultSpyFirestoreClient())
		response := httptest.NewRecorder()
		
		request := addContextFlagToRequest(newRecordVisitRequest(), DocumentSetFailFlag) 
		
		server.ServeHTTP(response, request)

		assertStatus(t, response.Result().StatusCode, http.StatusInternalServerError)
		assertResponseBodyLength(t, response.Body.Len(), 0)
	})
}

func assertResponseBodyLength(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got body length %d want %d", got, want)
	}
}