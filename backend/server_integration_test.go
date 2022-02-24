package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingVisitsAndGettingThem(t *testing.T) {
	server := &VisitCountServer{&InMemoryCounter{}}
	response := httptest.NewRecorder()

	server.ServeHTTP(httptest.NewRecorder(), newRecordVisitRequest())
	server.ServeHTTP(httptest.NewRecorder(), newRecordVisitRequest())
	server.ServeHTTP(httptest.NewRecorder(), newRecordVisitRequest())

	server.ServeHTTP(response, newGetVisitRequest())

	assertStatus(t, response.Result().StatusCode, http.StatusOK)
	assertVisitCount(t, response.Body.String(), "3")

}