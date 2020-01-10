package main

import (
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"
)

func TestEchoHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/v1/echo", nil)
	req.Header.Add("Accept-Encoding", "gzip")
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(EchoHandler)
	handler.ServeHTTP(rec, req)
	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		t.Fatal(err)
	}
	if rec.Body.String() != string(dump) {
		t.Errorf("body: got %v want %v", rec.Body.String(), string(dump))
	}
}
