// handlers_test.go
package main

import (
	"bytes"

	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestScraperHandler(t *testing.T) {

	data := url.Values{}
	data.Set("value", "150")
	b := bytes.NewBufferString(data.Encode())

	req, err := http.NewRequest("POST", "/scrape", b)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ScraperHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"shipHight":  150}`

	if strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
