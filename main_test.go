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

func createTestAlbum() album {
	testAlbum := album{ID: "10", Title: "testTitle", Artist: "testArtist", Price: 10.00}
	storage.Create(testAlbum)
	return testAlbum
}

func TestGetAlbums(t *testing.T) {
	req, _ := http.NewRequest("GET", "/albums", strings.NewReader(""))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatal("status must be 200")
	}
}

func TestGetOnePASS(t *testing.T) {
	testAlbum := createTestAlbum()
	req, _ := http.NewRequest("GET", "/albums/"+testAlbum.ID, strings.NewReader(""))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatal("status must be 200")
	}
}

func TestGetOneFAIL(t *testing.T) {
	requestPath := "1000"
	req, _ := http.NewRequest("GET", "/albums/"+requestPath, strings.NewReader(""))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatal("status must be 404")
	}
}

func TestDeleteOnePASS(t *testing.T) {
	testAlbum := createTestAlbum()
	req, _ := http.NewRequest("DELETE", "/albums/"+testAlbum.ID, strings.NewReader(""))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatal("status must be 204")
	}
}

func TestDeleteOneFAIL(t *testing.T) {
	requestPath := "1000"
	req, _ := http.NewRequest("DELETE", "/albums/"+requestPath, strings.NewReader(""))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatal("status must be 404")
	}
}

func TestUpdateOnePASS(t *testing.T) {
	testAlbum := createTestAlbum()
	req, _ := http.NewRequest("PUT", "/albums/"+testAlbum.ID, strings.NewReader(`{"title": "test"}`))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatal("status must be 200", w.Code)
	}
}

func TestUpdateOneFAIL(t *testing.T) {
	requestPath := "1001"
	req, _ := http.NewRequest("PUT", "/albums/"+requestPath, strings.NewReader(`{"title": "test"}`))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatal("status must be 404")
	}
}

func TestPostPASS(t *testing.T) {
	req, _ := http.NewRequest("POST", "/albums", strings.NewReader(`{"title": "test"}`))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Fatal("status must be 201", w.Code)
	}
}

func TestPostFAIL(t *testing.T) {
	req, _ := http.NewRequest("POST", "/albums", strings.NewReader(""))
	w := httptest.NewRecorder()
	requestHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatal("status must be 400", w.Code)
	}
}
