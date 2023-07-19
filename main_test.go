package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func requestHandler(w *httptest.ResponseRecorder, req *http.Request) {
	router := getRouter()
	router.ServeHTTP(w, req)
}

func TestGetAlbums(t *testing.T) {
	req, _ := http.NewRequest("GET", "/albums", strings.NewReader(""))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatal("status must be 200")
	}
}

func TestPostAlbumCorrectStructure(t *testing.T) {
	req, _ := http.NewRequest("POST", "/albums", strings.NewReader(`{"title": "test"}`))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Fatal("status must be 201", w.Code)
	}
}

func TestPostAlbumBadStructure(t *testing.T) {
	req, _ := http.NewRequest("POST", "/albums", strings.NewReader(""))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatal("status must be 400", w.Code)
	}
}

func TestAlbumDetail(t *testing.T) {
	req, _ := http.NewRequest("GET", "/albums/1", strings.NewReader(""))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatal("status must be 200")
	}
}

func TestAlbumNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/albums/1000", strings.NewReader(""))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatal("status must be 404")
	}
}

func TestDeleteAlbum(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/albums/1", strings.NewReader(""))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatal("status must be 204")
	}
}

func TestDeleteAlbumNotFound(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/albums/1000", strings.NewReader(""))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatal("status must be 404")
	}
}

func TestUpdateAlbum(t *testing.T) {
	req, _ := http.NewRequest("PUT", "/albums/2", strings.NewReader(`{"title": "test"}`))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatal("status must be 200", w.Code)
	}
}

func TestUpdateAlbumNotFound(t *testing.T) {
	req, _ := http.NewRequest("PUT", "/albums/1000", strings.NewReader(""))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatal("status must be 404")
	}
}
