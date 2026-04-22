package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Sem subir o servidor
func TestGetUsersWhitoutUpperServer(t *testing.T) {
	req := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(GetUsers)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status wait 200 and received: %d", w.Code)
	}
}

// subindo o servidor
func TestGetUsersUpServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(GetUsers))

	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 and found %d", resp.StatusCode)
	}
}
