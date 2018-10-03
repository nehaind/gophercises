package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

//executeRequest creates a new request and new recorder
func executeRequest(method string, url string, handler http.Handler) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	rr := httptest.NewRecorder()
	rr.Result()
	handler.ServeHTTP(rr, req)
	return rr, err
}
func TestMain(m *testing.M) {

	dashtest.ControlCoverage(m)
}

func TestSourceCodeHandler(t *testing.T) {
	rr, err := executeRequest("GET", "/debug?path=/home/neha/dev/src/gophercises_neha/exercise15/main.go", http.HandlerFunc(sourceCodeHandler))
	if err != nil {
		t.Fatal(err)
	}
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	rr, err = executeRequest("GET", "/debug?line=513&path=/home/neha/dev/src/gophercises_neha/exercise15/main.go", http.HandlerFunc(sourceCodeHandler))
	if err != nil {
		t.Fatal(err)
	}
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//invalid
	rr, _ = executeRequest("GET", "/debug?path=/home/neha/dev/src/gophercises_neha/exercise15/main1.go", http.HandlerFunc(sourceCodeHandler))

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}

func TestDevMw(t *testing.T) {
	handler := http.HandlerFunc(panicHandler)
	executeRequest("Get", "/panic", devMw(handler))

}

func TestIndex(t *testing.T) {

	req, err := http.NewRequest("GET", "/index", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestMainFunc(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("In Main panic")
		}
	}()
	go main()
	time.Sleep(1 * time.Second)
}
