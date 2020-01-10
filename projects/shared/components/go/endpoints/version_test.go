package endpoints

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewVersionHandler(t *testing.T) {
	version := "test"
	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(NewVersionHandler(version))
	handler.ServeHTTP(rec, req)
	expectedCode := http.StatusOK
	if status := rec.Code; status != expectedCode {
		t.Errorf("status code: got %v want %v", status, expectedCode)
	}
	if rec.Body.String() != version {
		t.Errorf("body: got %v want %v", rec.Body.String(), version)
	}
}
