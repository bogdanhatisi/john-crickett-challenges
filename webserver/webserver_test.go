package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleConnection(t *testing.T) {
	// Create a request to send to the server
	req, err := http.NewRequest("GET", "http://localhost/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler function to handle the request
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate the connection handling
		path := r.URL.Path
		response := "You requested the path: " + path
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(response))
	})

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "You requested the path: /"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

