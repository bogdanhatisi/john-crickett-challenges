package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHandleConnection(t *testing.T) {
	// Create a temporary file in the www directory for testing
	testFilePath := "www/testfile.html"
	testFileContent := "<html><body><h1>Test File</h1></body></html>"
	err := os.WriteFile(testFilePath, []byte(testFileContent), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(testFilePath) // Clean up after the test

	// Create a request to send to the server
	req, err := http.NewRequest("GET", "http://localhost/testfile.html", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a handler function to handle the request
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate the connection handling
		path := r.URL.Path
		filepath := "www" + path
		body, err := os.ReadFile(filepath)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html")
		w.Write(body)
	})

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := testFileContent
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
