package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)


func TestGetAllArticles(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/getAllArticles", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllArticles)

	handler.ServeHTTP(rr, req)


	// Check the response status code and content type
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Handler returned wrong content type: got %v want %v", contentType, "application/json")
	}


	var articles []Articles
	err = json.Unmarshal(rr.Body.Bytes(), &articles)
	if err != nil {
		t.Fatal(err)
	}


	if len(articles) < 1 {
		t.Error("Handler returned an empty list, expected at least one entry")
	}

	expectedType := reflect.TypeOf(Articles{})
	for _, article := range articles {
		actualType := reflect.TypeOf(article)
		if actualType != expectedType {
			t.Errorf("Handler returned an entry of unexpected type: got %v want %v", actualType, expectedType)
		}
	}

	log.Print(articles)
}