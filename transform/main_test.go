package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
func TestMainFunc(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Main panicked ??")
		}
	}()

	go main()
	time.Sleep(1 * time.Second)

}

func TestIndex(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
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

func TestTempFile(t *testing.T) {
	_, err := tempfile("", "png")
	if err != nil {
		t.Error("error found", err)
	}
}

// func TestUpload(t *testing.T) {

// 	body := &bytes.Buffer{}
// 	writer := multipart.NewWriter(body)

// 	req, _ := http.NewRequest("POST", "/upload", body)
// 	//r, _ := http.NewRequest("POST", srv.URL+"/upload", body)
// 	req.Header.Set("Content-Type", writer.FormDataContentType())
// 	res, _ := http.DefaultClient.Do(req)
// 	if res.StatusCode != http.StatusOK {
// 		t.Errorf("Expected status ok but got different status %v", res.Status)
// 	}

//wrong
// if err != nil {
// 	t.Fatal(err)
// }
// rr := httptest.NewRecorder()
// handler := http.HandlerFunc(upLoadHandler)
// handler.ServeHTTP(rr, req)

// if status := rr.Code; status != http.StatusOK {
// 	t.Errorf("handler returned wrong status code: got %v want %v",
// 		status, http.StatusOK)
// }

//}
