package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleRequests(t *testing.T) {
	// Create a test request to simulate incoming requests
	req, err := http.NewRequest("GET", "/SetCreateUsersRate?rate=0", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the HandleRequests function
	HandleRequests()

	// Send the test request to the server
	http.DefaultServeMux.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Add more assertions as needed
}
