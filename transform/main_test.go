package main

import (
	"bytes"
	"gophercises/Exercise_18/primitive"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/dash/dash_utils/dashtest"
)

func TestM(t *testing.T) {
	templistenAndServe := listenAndServeFunc
	defer func() {
		listenAndServeFunc = templistenAndServe
	}()
	listenAndServeFunc = func(port string, hanle http.Handler) error {
		panic("testing")
	}
	assert.PanicsWithValuef(t, "testing", main, "they should be equal")
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
func TestIndex(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	srv := httptest.NewServer(mux)
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	r, _ := http.NewRequest("GET", srv.URL, nil)
	res, _ := client.Do(r)
	if res.StatusCode != http.StatusOK {
		t.Error("Expected status ok but got different status")
	}
}

func TestUpload(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	srv := httptest.NewServer(mux)
	defer srv.Close()
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img/golang.png")
	file, err := os.Open(imgPath)
	if err != nil {
		t.Error("error in opening file")
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", file.Name())
	if err != nil {
		t.Error("error in copy")
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Error("error in copy")
	}
	err = writer.Close()
	if err != nil {
		t.Error("error in close writer")
	}
	r, _ := http.NewRequest("POST", srv.URL+"/upload", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	res, _ := http.DefaultClient.Do(r)
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status ok but got different status %v", res.Status)
	}
}

func TestCreateTempFile(t *testing.T) {
	_, err := createTempFile("/invalid/invalid", "txt")
	if err == nil {
		t.Error("Expected error but got no error")
	}
}

func TestModifyMode(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", modify)
	srv := httptest.NewServer(mux)
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/golang.png?mode=3", nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status ok but got different status %v", res.Status)
	}
}

func TestModifyModeNegative(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", modify)
	srv := httptest.NewServer(mux)
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/golang.png?mode=a", nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == http.StatusBadRequest {
		t.Errorf("excepted status: %v found: %v", 400, res.StatusCode)
	}
}

func TestModifyModeNegativeExt(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", modify)
	srv := httptest.NewServer(mux)
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/ghoper.txt?mode=2", nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == 400 {
		t.Errorf("excepted status: %v found: %v", http.StatusOK, res.Status)
	}
}

func TestModifyModeShapes(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", modify)
	srv := httptest.NewServer(mux)
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/ghoper.jpg?mode=3&n=5", nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		t.Errorf("excepted status: %v found: %v", http.StatusOK, res.Status)
	}
}

func TestGenImage(t *testing.T) {
	rs := bytes.NewReader(nil)
	mode := primitive.ModeCombo
	_, err := genImage(rs, "txt", -1, mode)
	if err == nil {
		t.Error("Expected error but no error")
	}
}
