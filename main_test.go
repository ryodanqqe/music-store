package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAlbums(t *testing.T) {
	rPath := "/albums"
	router := getRouter()
	req, _ := http.NewRequest("GET", rPath, strings.NewReader(""))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatal("status not 200")
	}
}

func TestAlbumDetail(t *testing.T) {
	rPath := "/albums/"
	router := getRouter()
	req, _ := http.NewRequest("GET", rPath+"1", strings.NewReader(""))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatal("status not 200")
	}
}

func TestAlbumNotFound(t *testing.T) {
	rPath := "/albums/"
	router := getRouter()
	req, _ := http.NewRequest("GET", rPath+"100", strings.NewReader(""))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Fatal("status not 404")
	}
}
